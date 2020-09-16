// Package logger ...
package logger

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

// Logger key constants definition.
const (
	ErrKey        = "_err"
	InfoKey       = "_info"
	WarnKey       = "_warn"
	CustomDataKey = "_custom_data"
)

// LogArr is array of string.
type LogArr []string

// Logger is the best logging lib.
type Logger interface {
	With(key string, value interface{}) Logger
	Error(format string, args ...interface{})
	Info(format string, args ...interface{})
	GetLogData(key string) LogArr
	Get(key string) interface{}
}

type logger struct {
	ctx context.Context
}

// GRPCLog logs with GRPC Context.
func GRPCLog(c context.Context) Logger {
	return &logger{
		ctx: c,
	}
}

func (_this *logger) With(key string, value interface{}) Logger {
	if value != nil {
		_this.Append(CustomDataKey, key)
		_this.Set(key, value)
	}
	return _this
}

func (_this *logger) Error(format string, args ...interface{}) {
	_this.Append(ErrKey, format, args...)
}

func (_this *logger) Info(format string, args ...interface{}) {
	_this.Append(InfoKey, format, args...)
}

func (_this *logger) Set(key string, value interface{}) {
	if _this.ctx != nil {
		_this.ctx = context.WithValue(_this.ctx, key, value)
	}
}

func (_this *logger) Get(key string) interface{} {
	if _this.ctx != nil {
		val := _this.ctx.Value(key)
		return val
	}
	return nil
}

func (_this *logger) GetLogData(key string) LogArr {
	val := _this.Get(key)
	if val != nil {
		return val.(LogArr)
	}
	return nil
}

func (_this *logger) Init(key string) {
	if _this.ctx != nil {
		val := _this.ctx.Value(key)
		if val == nil {
			_this.ctx = context.WithValue(_this.ctx, key, LogArr{})
		}
	}
}

func (_this *logger) Append(key string, format string, args ...interface{}) {
	_this.Init(key)
	if value := _this.Get(key); value != nil {
		logArr := value.(LogArr)
		logArr = append(logArr, fmt.Sprintf(format, args...))
		_this.Set(key, logArr)
	}
}

// Log function use zerolog package for write logs.
func Log(format string, args ...interface{}) {
	log.Log().Str("_msg", fmt.Sprintf(format, args...)).Msg("")
}
