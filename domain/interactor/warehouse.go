package interactor

import (
	"gintoki/domain/entity"
	"gintoki/domain/repository"
	"log"
)

type WarehouseInteractor interface {
	GetWarehouseCodeByID(warehouseID int64) (string, error)
}

type warehouseInteractor struct {
	warehouseRepository      repository.WarehouseRepository
	warehouseCacheRepository repository.WarehouseCacheRepository
}

func NewWarehouseInteractor(
	warehouseRepository repository.WarehouseRepository,
	warehouseCacheRepository repository.WarehouseCacheRepository) WarehouseInteractor {
	return &warehouseInteractor{
		warehouseRepository:      warehouseRepository,
		warehouseCacheRepository: warehouseCacheRepository,
	}
}

func (s *warehouseInteractor) GetWarehouseCodeByID(id int64) (string, error) {
	warehouseCode, err := s.warehouseCacheRepository.GetWarehouseCodeByWarehouseID(id)
	if err != nil {
		warehouseCode, err = s.warehouseRepository.GetWarehouseCodeByWarehouseID(id)
		if err != nil {
			return "", err
		}
		err = s.warehouseCacheRepository.UpdateWarehouseCodeByWarehouseID(entity.Warehouse{ID: id, Code: warehouseCode})
		if err != nil {
			log.Println("UpdateWarehouseCodeByWarehouseID failed:", err)
		}
	}
	return warehouseCode, err
}
