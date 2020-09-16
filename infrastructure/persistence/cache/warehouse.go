package cache

import (
	"encoding/json"
	"fmt"
	"gintoki/domain/entity"
	"gintoki/domain/repository"
	"gintoki/utils/localcache"
)

const (
	warehouseKeyTempl = "warehouse:%v"
)

type warehouseCacheRepository struct {
	cache localcache.LocalCache
}

func NewWarehouseCacheRepository(cache localcache.LocalCache) repository.WarehouseCacheRepository {
	return &warehouseCacheRepository{cache: cache}
}

func (s *warehouseCacheRepository) GetWarehouseCodeByWarehouseID(warehouseID int64) (string, error) {
	key := fmt.Sprintf(warehouseKeyTempl, warehouseID)
	data, err := s.cache.Get(key)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func (s *warehouseCacheRepository) UpdateWarehouseCodeByWarehouseID(warehouse entity.Warehouse) error {
	key := fmt.Sprintf(warehouseKeyTempl, warehouse.ID)
	data, err := json.Marshal(warehouse)
	if err != nil {
		return err
	}
	return s.cache.SetWithExpire(key, data, -1)
}
