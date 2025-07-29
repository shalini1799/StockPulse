package requestmodels

import (
	"time"
)

type Order struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	ProductID    int       `json:"product_id"`
	ProductName  string    `json:"product_name"`
	Quantity     int       `json:"quantity"`
	Status 	     string    `json:"status"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime`
	UpdatedAt	time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}