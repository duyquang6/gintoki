package database

import (
	"fmt"
	"gintoki/domain/repository"

	"github.com/jmoiron/sqlx"
)

type warehouseRepository struct {
	db *sqlx.DB
}

func (s *warehouseRepository) GetWarehouseCodeByWarehouseID(warehouseID int64) (string, error) {
	warehouseCode := ""
	err := s.db.Get(&warehouseCode, "SELECT code from ors_warehouse WHERE id=$1 LIMIT 1", warehouseID)
	fmt.Println(warehouseCode)
	return warehouseCode, err
}

func NewWarehouseRepository(db *sqlx.DB) repository.WarehouseRepository {
	return &warehouseRepository{
		db: db,
	}
}
