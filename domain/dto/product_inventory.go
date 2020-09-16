package dto

import "encoding/json"

type ProductInventory struct {
	ProductID           int64 `json:"product_id"`
	WarehousesInventory []WarehouseInventory
}

type WarehouseInventory struct {
	WarehouseCode string
	QtySalable    int
}

type KafkaMessage struct {
	Schema  json.RawMessage     `json:"-"`
	Payload KafkaConnectPayload `json:"payload"`
}

type KafkaConnectPayload struct {
	Before                  json.RawMessage         `json:"-"`
	UpdateLocalCacheRequest UpdateLocalCacheRequest `json:"after"`
	Source                  json.RawMessage         `json:"-"`
	TimestampMilisecond     uint64                  `json:"ts_ms"`
}

type UpdateLocalCacheRequest struct {
	WarehouseID int64  `json:"warehouse_id"`
	ProductID   int64  `json:"product_id"`
	UpdatedAt   string `json:"updated_at"`
	QtySalable  int    `json:"qty_salable"`
}
