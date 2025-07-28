// order-service/handlers/order.go
package handlers

import (
	"net/http"
	request_models "order-service/internal/models/request_models"
	response_models "order-service/internal/models/response_models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderService struct {
	DB *gorm.DB
}

func (handler *OrderService) CreateOrder(c *gin.Context) {
	tx := handler.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var order []request_models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(order) == 0 {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "No orders provided"})
		return
	}

	var ordersData []response_models.OrdersData
	var validOrders []request_models.Order

	for _, item := range order {
		available, stockData, err := checkStock(item.ProductID, tx)
		if err != nil {
			ordersData = append(ordersData, response_models.OrdersData{
				ProductID: item.ProductID,
				Error:     err.Error(),
			})
			continue
		}

		if available {
			ordersData = append(ordersData, response_models.OrdersData{
				ID:                item.ID,
				ProductID:         item.ProductID,
				Quantity:          item.Quantity,
				ProductName:       stockData.ProductName,
				InventoryID:       stockData.Id,
				InventoryQuantity: stockData.Quantity,
				Error:             "",
			})
			validOrders = append(validOrders, item)
		}
	}

	if len(validOrders) > 0 {
		if err := tx.Create(&validOrders).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create valid orders"})
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"message":    "Orders processing complete",
		"ordersData": ordersData,
	})
}


func checkStock(id int, db *gorm.DB) (bool, request_models.Inventory, error) {

	var inventory request_models.Inventory

	if err := db.Where("product_id = ?", id).Find(&inventory).Error; err != nil {
		return false, request_models.Inventory{}, err
	}

	return inventory.Quantity > 0, inventory, nil
}
	