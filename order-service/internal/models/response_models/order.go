package responsemodel

import(
	"time"
)

type OrdersData struct {
	ID                 int        `json:"id" gorm:"primaryKey"`
	ProductID          int        `json:"product_id"`
	Quantity           int        `json:"quantity"`
	CreatedAt          time.Time  `json:"created_at" gorm:"autoCreateTime"`
	ProductName        string     `json:"product_name"`
	InventoryID        int        `json:"inventory_id"`
	InventoryQuantity  int        `json:"inventory_quantity"`
	Error              string     `json:"error,omitempty"`
}
