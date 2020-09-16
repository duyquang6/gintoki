package middleware

import (
	"context"
	"fmt"
	"os"
	"time"

	"gintoki/config"
	"gintoki/utils/localcache"
	"gintoki/utils/logger"
	"gintoki/utils/loggerV2"

	"github.com/rs/zerolog/diode"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	wr = diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})
)

func GRPCLogging(cacheRepo localcache.LocalCache) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Before handler
		var (
			startTime = time.Now()
			method    = info.FullMethod
			msg       = "finish router"
		)
		log := zerolog.New(wr).With().Timestamp().Logger()
		zrLogger := log.With().
			Str("app_name", "gintoki").
			Str("version", config.AppConfig.App.Version).
			Str("endpoint", method).
			Str("req", fmt.Sprintf("%v", req)).
			Logger()

		res, err := handler(ctx, req)

		// After handler
		for _, key := range []string{logger.ErrKey, logger.InfoKey, logger.WarnKey} {
			zrLogger = loggerV2.PrepareLogData(ctx, zrLogger, key)
		}

		var (
			finishTime = time.Now()
			latency    = finishTime.Sub(startTime)
			statusCode = codes.OK
		)
		if err != nil {
			statusCode = codes.Unknown
		}
		if config.AppConfig.EnableCacheStatistics {
			zrLogger = zrLogger.With().Dict("cache_statistics", zerolog.Dict().Fields(cacheRepo.GetStatistics())).Logger()
		}
		zrLogger.Log().
			Dur("latency", latency).
			Int("status_code", int(statusCode)).
			Msg(msg)
		loggerV2.CleanLogData(ctx)
		return res, err
	}
}
