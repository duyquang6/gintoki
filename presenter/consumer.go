package presenter

import (
	"context"
	"fmt"
	"gintoki/application/handler"
	"gintoki/config"
	"gintoki/utils/logger"
	"gintoki/utils/loggerV2"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

type kafkaConsumer struct {
	kafkaClient sarama.ConsumerGroup
	cancelCtx   context.CancelFunc
	consumer    *Consumer
}

var (
	wr = diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})
)

func (s *kafkaConsumer) Run() {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	appKafkaConfig := config.AppConfig.Kafka
	config := sarama.NewConfig()
	switch appKafkaConfig.Assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	default:
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	}
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	version, err := sarama.ParseKafkaVersion(appKafkaConfig.Version)
	config.Version = version
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	s.cancelCtx = cancel
	brokers := strings.Split(appKafkaConfig.Server, ",")

	s.kafkaClient, err = sarama.NewConsumerGroup(brokers, uuid.New().String(), config)
	go func() {
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := s.kafkaClient.Consume(ctx, strings.Split(appKafkaConfig.Topic, ","), s.consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			s.consumer.ready = make(chan bool)
		}
	}()
	<-s.consumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")
}

func (s *kafkaConsumer) Close() {
	s.cancelCtx()
	if err := s.kafkaClient.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready                   chan bool
	ProductInventoryHandler handler.ProductInventoryHandler
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		ctx := context.WithValue(context.Background(), "request-id", uuid.New().String())
		var (
			startTime = time.Now()
		)
		log := zerolog.New(wr).With().Timestamp().Logger()
		zrLogger := log.With().
			Str("app_name", "gintoki_consumer").
			Str("version", config.AppConfig.App.Version).
			Logger()

		err := consumer.ProductInventoryHandler.UpdateLocalCacheFromQueue(ctx, message.Value)
		session.MarkMessage(message, "")

		// After handler
		for _, key := range []string{logger.ErrKey, logger.InfoKey, logger.WarnKey} {
			zrLogger = loggerV2.PrepareLogData(ctx, zrLogger, key)
		}
		var (
			finishTime = time.Now()
			latency    = finishTime.Sub(startTime)
			statusCode = http.StatusOK
		)
		if err != nil {
			statusCode = http.StatusBadRequest
		}
		zrLogger.Log().
			Dur("latency", latency).
			Int("status_code", int(statusCode)).
			Msg("finish router")
		loggerV2.CleanLogData(ctx)
		return err
	}
	return nil
}

func NewConsumer(productInventoryHandler handler.ProductInventoryHandler) Server {
	server := kafkaConsumer{
		consumer: &Consumer{
			ready:                   make(chan bool),
			ProductInventoryHandler: productInventoryHandler,
		},
	}
	return &server
}
