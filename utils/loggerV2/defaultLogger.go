package loggerV2

import (
	"context"

	"github.com/rs/zerolog"
)

var defaultLogger LoggerV2

func init() {
	defaultLogger = newLogger()
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Infof(ctx, format, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Warnf(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Errorf(ctx, format, args...)
}

func GetLogData(ctx context.Context, key string) []string {
	return defaultLogger.GetLogData(ctx, key)
}

func CleanLogData(ctx context.Context) {
	defaultLogger.CleanLogData(ctx)
}

func PrepareLogData(ctx context.Context, zrLogger zerolog.Logger, key string) zerolog.Logger {
	logData := GetLogData(ctx, key)
	zrLogger = zrLogger.With().Strs(key, logData).Logger()
	return zrLogger
}
