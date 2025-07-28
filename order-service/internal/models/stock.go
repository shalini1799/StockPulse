// order-service/models/order.go
package models

import (
	"time"
)

type Order struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	ProductID  uint      `json:"product_id"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
}