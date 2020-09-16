package presenter

import (
	"fmt"
	"gintoki/application/handler"
	pb "gintoki/application/handler/proto"
	"gintoki/application/middleware"
	"gintoki/config"
	"gintoki/utils/localcache"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	server *grpc.Server
}

func getGRPCServer(cacheRepo localcache.LocalCache, productInventoryHandler handler.ProductInventoryHandler) *grpc.Server {

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_prometheus.UnaryServerInterceptor,
			middleware.MetricInterceptor(cacheRepo),
			middleware.AddRequestID(),
			middleware.GRPCLogging(cacheRepo),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
				return status.Errorf(codes.Unknown, "panic triggered: %v", p)
			})),
		),
	)
	pb.RegisterProductInventoryServiceServer(server, productInventoryHandler)
	// Register reflection service on gRPC server.
	reflection.Register(server)

	// Register prometheus
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(server)
	return server
}

func (s *grpcServer) Run() {
	appConfig := config.AppConfig.App
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", appConfig.Port))
	if err != nil {
		panic(err)
	}
	// start gRPC server
	log.Println("Starting listen gRPC server on", appConfig.Port)
	if err := s.server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

func (s *grpcServer) Close() {
	log.Println("shutting down gRPC server...")
	s.server.GracefulStop()
}

func NewGRPCServer(cacheRepo localcache.LocalCache, productInventoryHandler handler.ProductInventoryHandler) Server {
	server := getGRPCServer(cacheRepo, productInventoryHandler)
	return &grpcServer{server}
}
