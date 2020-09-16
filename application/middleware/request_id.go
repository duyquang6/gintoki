package middleware

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func AddRequestID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = context.WithValue(ctx, "request-id", uuid.New().String())
		return handler(ctx, req)
	}
}
