package interactor

import (
	"context"
	pb "gintoki/application/handler/proto"
	"gintoki/config"
	"gintoki/domain/dto"
	"gintoki/domain/repository"
	"gintoki/domain/tool"
	"gintoki/utils/loggerV2"
	"log"

	"github.com/allegro/bigcache"
	"github.com/jinzhu/copier"
)

type ProductInventoryInteractor interface {
	Get(id int) (*dto.ProductInventory, error)
	GetMulti(ctx context.Context, productIDs []int64) ([]*pb.GetMultiProductInventoryData, error)
	UpdateLocalCacheFromQueue(ctx context.Context, req dto.UpdateLocalCacheRequest) error
}

type productInventoryInteractor struct {
	productInventoryRepository repository.ProductInventoryRepository
	productInventoryCacheRepo  repository.ProductInventoryCacheRepo
	warehouseInteractor        WarehouseInteractor
	idGenerator                tool.IdGenerator
}

func NewProductInventoryInteractor(
	productInventoryRepository repository.ProductInventoryRepository,
	productInventoryCacheRepo repository.ProductInventoryCacheRepo,
	warehouseInteractor WarehouseInteractor,
	idGenerator tool.IdGenerator,
) ProductInventoryInteractor {
	return &productInventoryInteractor{
		productInventoryRepository: productInventoryRepository,
		productInventoryCacheRepo:  productInventoryCacheRepo,
		warehouseInteractor:        warehouseInteractor,
		idGenerator:                idGenerator,
	}
}

func (s *productInventoryInteractor) Get(productID int) (*dto.ProductInventory, error) {
	res := &dto.ProductInventory{}
	product, err := s.productInventoryCacheRepo.GetByProductID(int64(productID))
	if err == nil {
		_ = copier.Copy(res, product)
		return res, err
	}
	if err != bigcache.ErrEntryNotFound {
		return nil, err
	}
	product, err = s.productInventoryRepository.GetByProductID(int64(productID))
	if err != nil {
		return nil, err
	}
	err = s.productInventoryCacheRepo.UpdateIfNotExist(*product)
	if err != nil {
		log.Println("cannot update if not exist", err)
	}
	_ = copier.Copy(res, product)
	return res, nil
}

func (s *productInventoryInteractor) GetMulti(ctx context.Context, productIDs []int64) ([]*pb.GetMultiProductInventoryData, error) {
	var res []*pb.GetMultiProductInventoryData
	isUseCache := config.AppConfig.UseCache
	var productIDsNotFound []int64
	if isUseCache {
		for _, productID := range productIDs {
			productEntity, err := s.productInventoryCacheRepo.GetByProductID(productID)
			if err != nil {
				productIDsNotFound = append(productIDsNotFound, productID)
				continue
			}
			product := pb.GetMultiProductInventoryData{}
			err = copier.Copy(&product, productEntity)
			err = copier.Copy(&product.Data, productEntity.WarehousesInventory)
			if err != nil {
				loggerV2.Errorf(ctx, "copy failed:%v", err)
				productIDsNotFound = append(productIDsNotFound, productID)
				continue
			}
			res = append(res, &product)
		}

		if len(productIDsNotFound) == 0 {
			return res, nil
		}
	} else {
		productIDsNotFound = productIDs
	}

	productsInventory, err := s.productInventoryRepository.GetMultiByProductID(productIDsNotFound)
	if err != nil {
		loggerV2.Errorf(ctx, "Error when call GetMultiByProductID, err: %s", err.Error())
		return nil, err
	}

	if isUseCache {
		for _, v := range productsInventory {
			err = s.productInventoryCacheRepo.UpdateIfNotExist(v)
			if err != nil {
				loggerV2.Errorf(ctx, "UpdateIfNotExist product inventory failed: %s", err.Error())
			}
		}
	}

	for _, p := range productsInventory {
		product := pb.GetMultiProductInventoryData{}
		err = copier.Copy(&product, p)
		err = copier.Copy(&product.Data, p.WarehousesInventory)
		if err != nil {
			loggerV2.Errorf(ctx, "Error when copy data, err: %s", err.Error())
			continue
		}
		res = append(res, &product)
	}
	return res, nil
}

func (s *productInventoryInteractor) UpdateLocalCacheFromQueue(ctx context.Context, req dto.UpdateLocalCacheRequest) error {
	warehouseCode, err := s.warehouseInteractor.GetWarehouseCodeByID(req.WarehouseID)
	if err != nil {
		loggerV2.Errorf(ctx, "cannot get warehouse code from id %v, err %v", req.WarehouseID, err)
		return err
	}
	err = s.productInventoryCacheRepo.UpdateByProductIDAndWarehouseCode(req.ProductID, warehouseCode, req.QtySalable)
	if err != nil {
		loggerV2.Errorf(ctx, "UpdateLocalCacheFromQueue failed for product_id %v, warehouse_code %v, err:%v", req.ProductID, warehouseCode, err)
		return err
	}
	loggerV2.Infof(ctx, "Update cache success, productID:%v, warehouseCode:%v, qty_salable:%v", req.ProductID, warehouseCode, req.QtySalable)
	return nil
}
