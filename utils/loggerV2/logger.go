package loggerV2

import (
	"context"
	"fmt"
	"gintoki/utils/logger"
	"gintoki/utils/request_id"
)

var (
	errorsLogMap = make(map[string][]string)
	warnsLogMap  = make(map[string][]string)
	infosLogMap  = make(map[string][]string)
)

type LoggerV2 interface {
	Infof(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	GetLogData(ctx context.Context, key string) []string
	CleanLogData(ctx context.Context)
}
type loggerV2 struct{}

func newLogger() LoggerV2 {
	return &loggerV2{}
}

func (l *loggerV2) Infof(ctx context.Context, format string, args ...interface{}) {
	if key := request_id.GetRequestIDFromCtx(ctx); len(key) > 0 {
		infosLogMap[key] = append(infosLogMap[key], fmt.Sprintf(format, args...))
	}
}

func (l *loggerV2) Warnf(ctx context.Context, format string, args ...interface{}) {
	if key := request_id.GetRequestIDFromCtx(ctx); len(key) > 0 {
		warnsLogMap[key] = append(warnsLogMap[key], fmt.Sprintf(format, args...))
	}
}

func (l *loggerV2) Errorf(ctx context.Context, format string, args ...interface{}) {
	if key := request_id.GetRequestIDFromCtx(ctx); len(key) > 0 {
		errorsLogMap[key] = append(errorsLogMap[key], fmt.Sprintf(format, args...))
	}
}

func (l *loggerV2) GetLogData(ctx context.Context, key string) []string {
	var data []string
	reqID := request_id.GetRequestIDFromCtx(ctx)
	switch key {
	case logger.ErrKey:
		return errorsLogMap[reqID]
	case logger.InfoKey:
		return infosLogMap[reqID]
	case logger.WarnKey:
		return warnsLogMap[reqID]
	}
	return data
}

func (l *loggerV2) CleanLogData(ctx context.Context) {
	reqID := request_id.GetRequestIDFromCtx(ctx)
	delete(errorsLogMap, reqID)
	delete(infosLogMap, reqID)
	delete(warnsLogMap, reqID)
}
