package requestmodels

import (
	"time"
)

type Inventory struct {
	Id           int 	   `json:"id" gorm:"primaryKey"`
	ProductID    string    `json:"product_id" gorm:"uniqueIndex"`
	ProductName  string    `json:"product_name"`
	Quantity     int       `json:"quantity"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}