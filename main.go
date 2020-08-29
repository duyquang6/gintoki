package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"gintoki/database"
	"github.com/Shopify/sarama"
	"github.com/allegro/bigcache"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

var (
	cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
)


func main() {
	addr := flag.String("addr", ":8080", "server address")
	database.NewDB()
	//cache.Delete("my-unique-key")
	r := gin.Default()
	stockRep := database.NewStockRepo()
	r.GET("/cache", func(c *gin.Context) {
		productID, err :=strconv.Atoi( c.Query("product_id"))
		if err != nil {
			c.String(400, err.Error())
			return
		}

		entry, err := cache.Get("product:" + strconv.Itoa(productID))
		if err == nil {
			stockItem := database.StockItem{}
			fmt.Println("hit cache")
			json.NewDecoder(bytes.NewReader(entry)).Decode(&stockItem)
			c.String(200, "%v",stockItem)
			return
		}
		stockItem, err := stockRep.GetStockItem(productID)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		reqBodyBytes := new(bytes.Buffer)
		err = json.NewEncoder(reqBodyBytes).Encode(stockItem)
		err = cache.Set("product:" + strconv.Itoa(productID), reqBodyBytes.Bytes())
		if err != nil {
			fmt.Println(err)
		}
		c.String(200, "%v", stockItem)
	})

	consumeAndInvalidCache()

	r.Run(*addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type KafkaValuePayload struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	Before map[string]interface{} `json:"before"`
}

func consumeAndInvalidCache() {
	config := sarama.NewConfig()
	config.ClientID = "go-kafka-consumer-2"
	config.Consumer.Return.Errors = true
	brokers := []string{"192.168.1.104:9092"}
	// Create new consumer
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	consumer, errors := consume([]string{"dbserver1.inventory.tala_warehouse_stock_item"}, master)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many message processed
	msgCount := 0

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case msg:= <-consumer:
				msgCount++
				//fmt.Println("key:", string(msg.Key), "value:", string(msg.Value))
				kkPayload := KafkaValuePayload{}
				err = json.Unmarshal(msg.Value, &kkPayload)
				fmt.Println(kkPayload.Payload.Before)
				err = cache.Delete("product:" + strconv.Itoa(int(kkPayload.Payload.Before["product_id"].(float64))))
				if err != nil {
					fmt.Println(err)
				}
			case consumerError := <-errors:
				msgCount++
				fmt.Println("Received consumerError ", string(consumerError.Topic), string(consumerError.Partition), consumerError.Err)
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
}

func consume(topics []string, master sarama.Consumer) (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)
	for _, topic := range topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}
		partitions, _ := master.Partitions(topic)
		// this only consumes partition no 1, you would probably want to consume all partitions
		consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest)
		if nil != err {
			fmt.Printf("Topic %v Partitions: %v", topic, partitions)
			panic(err)
		}
		fmt.Println(" Start consuming topic ", topic)
		go func(topic string, consumer sarama.PartitionConsumer) {
			for {
				select {
				case consumerError := <-consumer.Errors():
					errors <- consumerError
					fmt.Println("consumerError: ", consumerError.Err)
				case msg := <-consumer.Messages():
					consumers <- msg
					//fmt.Println("Got message on topic ", topic, msg.Value)
				}
			}
		}(topic, consumer)
	}

	return consumers, errors
}
