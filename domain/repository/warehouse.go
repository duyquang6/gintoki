package repository

type WarehouseRepository interface {
	GetWarehouseCodeByWarehouseID(warehouseID int64) (string, error)
}
