package main

import (
	"fmt"
	"gintoki/application/handler"
	"gintoki/config"
	"gintoki/dependency"
	"gintoki/domain/interactor"
	"gintoki/presenter"
	"gintoki/utils/localcache"
	"gintoki/utils/multilock"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/coocood/freecache"
	_ "github.com/lib/pq"
)

var (
	cacheSize = 100 * 1024 * 1024
	cacheCore = freecache.NewCache(cacheSize)
	cacheRepo = localcache.NewLocalCacheRepository(cacheCore)
)

func main() {
	debug.SetGCPercent(20)

	config.InitConfig()

	// DB init
	db, err := dependency.NewPostgresSQLConnection()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	defer dependency.Close(db)

	multilocker := multilock.NewMultilock()

	productInventoryRepo := dependency.NewProductInventoryRepository(db)
	productInventoryCacheRepo := dependency.NewProductInventoryCacheRepo(cacheRepo, multilocker, productInventoryRepo)
	warehouseRepo := dependency.NewWarehouseRepo(db)
	warehouseCacheRepo := dependency.NewWarehouseCacheRepo(cacheRepo)

	warehouseInteractor := interactor.NewWarehouseInteractor(warehouseRepo, warehouseCacheRepo)

	productInventoryInteractor := interactor.NewProductInventoryInteractor(
		productInventoryRepo,
		productInventoryCacheRepo,
		warehouseInteractor,
		dependency.NewIdGenerator(),
	)
	productInventoryHandler := handler.ProductInventoryHandler{Interactor: productInventoryInteractor}

	// GRPC server
	grpcServer := presenter.NewGRPCServer(cacheRepo, productInventoryHandler)
	go grpcServer.Run()
	defer grpcServer.Close()

	// HTTP-GRPC proxy
	httpGRPCProxy := presenter.NewHTTPGRPCServer()
	go httpGRPCProxy.Run()
	defer httpGRPCProxy.Close()

	// Kafka consumer
	consumer := presenter.NewConsumer(productInventoryHandler)
	go consumer.Run()
	defer consumer.Close()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
}
