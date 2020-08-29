package database

type StockRepo interface {
	GetStockItem() (StockItem, error)
}

type stockRepo struct{}

func NewStockRepo() stockRepo {
	s:= stockRepo{}
	return s
}
type StockItem struct {
	ProductID  int    `db:"product_id"`
	SKU        string `db:"sku"`
	Name       string `db:"name"`
	QtySalable string `db:"qty_salable"`
}

func (s *stockRepo) GetStockItem(productID int) (StockItem, error) {
	tx := NewDB()
	stock := StockItem{}
	err := tx.Get(&stock,
		"SELECT product_id, sku, name, qty_salable FROM tala_warehouse_stock_item WHERE product_id=? LIMIT 1", productID)
	return stock, err
}
