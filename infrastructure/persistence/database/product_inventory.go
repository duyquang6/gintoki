package database

import (
	"gintoki/domain/entity"
	"gintoki/domain/repository"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type productInventoryRepository struct {
	db *sqlx.DB
}

type dataSQL struct {
	ProductID      int64  `db:"product_id"`
	WarehouseCodes string `db:"warehouse_codes"`
	QtySalables    string `db:"qty_salable"`
}

func (s *productInventoryRepository) GetMultiByProductID(productIDs []int64) ([]entity.ProductInventory, error) {
	data := []dataSQL{}
	products := []entity.ProductInventory{}
	query, args, err := sqlx.In(`select product_id, array_to_string(array_agg(tw.code), ',') warehouse_codes, array_to_string(array_agg(qty_salable), ',') qty_salable
	from tala_warehouse_stock_item tl
		left join ors_warehouse tw on tw.id = tl.warehouse_id
	where tl.product_id in (?)
		and tw.is_active = 1
	group by tl.product_id;`, productIDs)
	if err != nil {
		return nil, err
	}
	query = s.db.Rebind(query)
	err = s.db.Select(&data, query, args...)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		warehouseCodes := strings.Split(v.WarehouseCodes, ",")
		qtySalables := strings.Split(v.QtySalables, ",")
		if len(qtySalables) != len(warehouseCodes) {
			log.Println("length of qty_salables and warehouse_codes different")
			continue
		}
		wi := []entity.WarehouseInventory{}
		for i := range warehouseCodes {
			qty, err := strconv.Atoi(qtySalables[i])
			if err != nil {
				qty = 0
			}
			wi = append(wi, entity.WarehouseInventory{
				QtySalable:    qty,
				WarehouseCode: warehouseCodes[i],
			})
		}
		product := entity.ProductInventory{
			ProductID:           v.ProductID,
			WarehousesInventory: wi,
		}
		products = append(products, product)
	}
	return products, err
}

func (s *productInventoryRepository) GetMockData(productIDs []int64) ([]entity.ProductInventory, error) {
	data := []dataSQL{}
	products := []entity.ProductInventory{}
	err := s.db.Select(&data,
		`select product_id, array_to_string(array_agg(tw.code), ',') warehouse_codes, array_to_string(array_agg(qty_salable), ',') qty_salable
		from tala_warehouse_stock_item tl
			left join ors_warehouse tw on tw.id = tl.warehouse_id
		where tw.is_active = 1
		group by tl.product_id limit 100000;`)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		warehouseCodes := strings.Split(v.WarehouseCodes, ",")
		qtySalables := strings.Split(v.QtySalables, ",")
		if len(qtySalables) != len(warehouseCodes) {
			log.Println("length of qty_salables and warehouse_codes different")
			continue
		}
		wi := []entity.WarehouseInventory{}
		for i := range warehouseCodes {
			qty, err := strconv.Atoi(qtySalables[i])
			if err != nil {
				qty = 0
			}
			wi = append(wi, entity.WarehouseInventory{
				QtySalable:    qty,
				WarehouseCode: warehouseCodes[i],
			})
		}
		product := entity.ProductInventory{
			ProductID:           v.ProductID,
			WarehousesInventory: wi,
		}
		products = append(products, product)
	}
	return products, err
}

func (s *productInventoryRepository) GetByProductID(productID int64) (*entity.ProductInventory, error) {
	data := dataSQL{}
	err := s.db.Get(&data,
		`select product_id, array_agg(tw.code) warehouse_codes, array_agg(qty_salable) qty_salable
		from tala_warehouse_stock_item tl
			left join ors_warehouse tw on tw.id = tl.warehouse_id
		where tl.product_id=$1
			and tw.is_active = 1
		group by tl.product_id limit 1;`, productID)
	wi := []entity.WarehouseInventory{}
	warehouseCodes := strings.Split(data.WarehouseCodes, ",")
	qtySalables := strings.Split(data.QtySalables, ",")
	for i := range warehouseCodes {
		qty, err := strconv.Atoi(qtySalables[i])
		if err != nil {
			qty = 0
		}
		wi = append(wi, entity.WarehouseInventory{
			QtySalable:    qty,
			WarehouseCode: warehouseCodes[i],
		})
	}
	product := &entity.ProductInventory{
		ProductID:           data.ProductID,
		WarehousesInventory: wi,
	}
	return product, err
}

func NewProductInventoryRepository(db *sqlx.DB) repository.ProductInventoryRepository {
	return &productInventoryRepository{
		db: db,
	}
}
