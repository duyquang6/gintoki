package dependency

import (
	"gintoki/domain/repository"
	"gintoki/infrastructure/persistence/cache"
	"gintoki/infrastructure/persistence/database"
	"gintoki/utils/localcache"
	"gintoki/utils/multilock"

	"github.com/jmoiron/sqlx"
)

func NewProductInventoryRepository(db interface{}) repository.ProductInventoryRepository {
	switch connection := db.(type) {
	case *sqlx.DB:
		return database.NewProductInventoryRepository(connection)
	}

	return nil
}

func NewProductInventoryCacheRepo(db interface{}, multilock multilock.Multilock, productInventoryRepo repository.ProductInventoryRepository) repository.ProductInventoryCacheRepo {
	switch connection := db.(type) {
	case localcache.LocalCache:
		return cache.NewProductInventoryCacheRepo(connection, multilock, productInventoryRepo)
	}
	return nil
}

func NewWarehouseRepo(db interface{}) repository.WarehouseRepository {
	switch connection := db.(type) {
	case *sqlx.DB:
		return database.NewWarehouseRepository(connection)
	}
	return nil
}

func NewWarehouseCacheRepo(db interface{}) repository.WarehouseCacheRepository {
	switch connection := db.(type) {
	case localcache.LocalCache:
		return cache.NewWarehouseCacheRepository(connection)
	}
	return nil
}
