// order-service/handlers/order.go
package handlers

import (
	"net/http"
	"order-service/constant"
	request_models "order-service/internal/models/request_models"
	response_models "order-service/internal/models/response_models"
	"strconv"

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
				ProductName: item.ProductName,
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

func (handler *OrderService) GetAllOrders(c *gin.Context) {
	var orders []request_models.Order
	if err := handler.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch orders"})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No orders found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func(handler *OrderService) UpdateOrder(c *gin.Context) {

	tx := handler.DB.Begin()

	var updatedOrder request_models.Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check order status
	if err := handler.DB.First(&updatedOrder, updatedOrder.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	if updatedOrder.Status == constant.OrderStatusShipped {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order cannot be updated if it is shipped"})
		return
	}

	if updatedOrder.Status == constant.OrderStatusPending {
		tx.Model(&request_models.Order{}).Where("id = ?", updatedOrder.ID).Updates(updatedOrder)
		if err := tx.Error; err != nil {	
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update order"})
			return
		}

		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully", "order": updatedOrder})
	}
	
}

func(handler *OrderService) DeleteOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	tx := handler.DB.Begin()

	//check order status
	var updatedOrder request_models.Order
	if err := tx.First(&updatedOrder, orderID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if updatedOrder.Status == constant.OrderStatusShipped {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order cannot be cancelled once if it is shipped"})
		return
	}

	if updatedOrder.Status == constant.OrderStatusPending {
		if err := tx.Model(&request_models.Order{}).Where("id = ?", orderID).Update("status", "cancelled").Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel order"})
			return
		}
 	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}

func (handler *OrderService) GetOrderByID(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order request_models.Order
	if err := handler.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}	

func checkStock(id int, db *gorm.DB) (bool, request_models.Inventory, error) {

	var inventory request_models.Inventory

	if err := db.Where("product_id = ?", id).Find(&inventory).Error; err != nil {
		return false, request_models.Inventory{}, err
	}

	return inventory.Quantity > 0, inventory, nil
}
	