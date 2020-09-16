package entity

type ProductInventory struct {
	ProductID           int64                `json:"product_id"`
	WarehousesInventory []WarehouseInventory `json:"warehouses_inventory"`
}
type WarehouseInventory struct {
	WarehouseCode string `json:"warehouse_code"`
	QtySalable    int    `json:"qty_salable"`
}
