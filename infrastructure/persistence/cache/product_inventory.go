package cache

import (
	"encoding/json"
	"fmt"
	"gintoki/domain/entity"
	"gintoki/domain/repository"
	"gintoki/utils/localcache"
	"gintoki/utils/multilock"

	"github.com/coocood/freecache"
)

type productInventoryCacheRepo struct {
	cache                localcache.LocalCache
	multilock            multilock.Multilock
	productInventoryRepo repository.ProductInventoryRepository
}

const (
	productInventoryKeyTempl = "product_inventory:%v"
)

func NewProductInventoryCacheRepo(
	cache localcache.LocalCache,
	multilock multilock.Multilock,
	productInventoryRepo repository.ProductInventoryRepository) repository.ProductInventoryCacheRepo {
	return &productInventoryCacheRepo{
		cache:                cache,
		multilock:            multilock,
		productInventoryRepo: productInventoryRepo,
	}
}

func (s *productInventoryCacheRepo) GetByProductID(productID int64) (*entity.ProductInventory, error) {
	key := fmt.Sprintf(productInventoryKeyTempl, productID)
	l := s.multilock.Get(key)
	l.RLock()
	defer func() {
		l.RUnlock()
		s.multilock.Put(key)
	}()
	product := entity.ProductInventory{}
	data, err := s.cache.Get(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *productInventoryCacheRepo) UpdateIfNotExist(product entity.ProductInventory) error {
	key := fmt.Sprintf(productInventoryKeyTempl, product.ProductID)
	l := s.multilock.Get(key)
	l.Lock()
	defer func() {
		l.Unlock()
		s.multilock.Put(key)
	}()
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}
	_, err = s.cache.GetOrSetWithExpire(key, data, -1)
	return err
}

func (s *productInventoryCacheRepo) Update(product entity.ProductInventory) error {
	key := fmt.Sprintf(productInventoryKeyTempl, product.ProductID)
	l := s.multilock.Get(key)
	l.Lock()
	defer func() {
		l.Unlock()
		s.multilock.Put(key)
	}()
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}
	return s.cache.SetWithExpire(fmt.Sprintf(productInventoryKeyTempl, product.ProductID), data, -1)
}

func (s *productInventoryCacheRepo) GetProductInventoryKey(id int64) string {
	return fmt.Sprintf(productInventoryKeyTempl, id)
}

func (s *productInventoryCacheRepo) UpdateByProductIDAndWarehouseCode(
	productID int64, warehouseCode string, qtySalable int) error {
	key := fmt.Sprintf(productInventoryKeyTempl, productID)
	l := s.multilock.Get(key)
	l.Lock()
	defer func() {
		l.Unlock()
		s.multilock.Put(key)
	}()

	product := &entity.ProductInventory{}
	data, err := s.cache.Get(key)
	_ = json.Unmarshal(data, product)
	if err != nil {
		if err != freecache.ErrNotFound {
			return err
		}
		// No entry found for this productID in cache
		product, err = s.productInventoryRepo.GetByProductID(productID)
		if err != nil {
			return err
		}
	}

	for _, warehouseInventory := range product.WarehousesInventory {
		if warehouseCode == warehouseInventory.WarehouseCode {
			warehouseInventory.QtySalable = qtySalable
		}
	}

	data, err = json.Marshal(product)
	if err != nil {
		return err
	}
	return s.cache.SetWithExpire(fmt.Sprintf(productInventoryKeyTempl, product.ProductID), data, -1)
}
