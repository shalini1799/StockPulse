package requestmodels

import (
	"time"
)

type Order struct {
	ID         int      `json:"id" gorm:"primaryKey"`
	ProductID  int      `json:"product_id"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
}