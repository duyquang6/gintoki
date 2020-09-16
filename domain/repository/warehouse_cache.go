package repository

import "gintoki/domain/entity"

type WarehouseCacheRepository interface {
	GetWarehouseCodeByWarehouseID(warehouseID int64) (string, error)
	UpdateWarehouseCodeByWarehouseID(warehouse entity.Warehouse) error
}
