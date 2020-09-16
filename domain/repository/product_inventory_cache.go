package repository

import "gintoki/domain/entity"

type ProductInventoryCacheRepo interface {
	GetByProductID(productID int64) (*entity.ProductInventory, error)
	UpdateIfNotExist(product entity.ProductInventory) error
	Update(product entity.ProductInventory) error
	UpdateByProductIDAndWarehouseCode(productID int64, warehouseCode string, qtySalable int) error
	GetProductInventoryKey(id int64) string
}
