package repository

import "gintoki/domain/entity"

type ProductInventoryRepository interface {
	GetByProductID(productID int64) (*entity.ProductInventory, error)
	GetMultiByProductID(productIDs []int64) ([]entity.ProductInventory, error)
	GetMockData(productIDs []int64) ([]entity.ProductInventory, error)
}
