package presenter

import (
	"context"
	pb "gintoki/application/handler/proto"
	"gintoki/config"
	"log"
	"net/http"
	"time"

	"net/http/pprof"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

type httpGRPCServer struct {
	srv *http.Server
}

func (s *httpGRPCServer) Run() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterProductInventoryServiceHandlerFromEndpoint(ctx, gwmux, ":8080", opts); err != nil {
		return
	}
	mux := http.NewServeMux()
	if config.AppConfig.App.EnablePPROF {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
		mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
		mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	}
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", gwmux)
	s.srv.Handler = mux
	// Run our server in a goroutine so that it doesn't block.
	log.Println("start http-grpc reverse proxy on port 6060")
	if err := s.srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (s *httpGRPCServer) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Println("err when close grpc-http proxy:", err)
	}
}

func NewHTTPGRPCServer() *httpGRPCServer {
	return &httpGRPCServer{
		srv: &http.Server{
			Addr:         ":6060",
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		},
	}
}
