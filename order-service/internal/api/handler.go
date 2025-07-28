// order-service/handlers/order.go
package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"order-service/internal/models"
	"order-service/internal/db"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order_id": order.ID})
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := db.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}