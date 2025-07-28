package api

import (
	"net/http"
	"stockpulse/internal/model"
	"stockpulse/internal/repository"
	"stockpulse/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	repo := repository.NewStockRepository()
	svc := service.NewStockService(repo)

	router.POST("/stocks", func(c *gin.Context) {
		var stock model.Stock
		if err := c.ShouldBindJSON(&stock); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		svc.AddStock(stock)
		c.JSON(http.StatusOK, gin.H{"message": "Stock added"})
	})

	router.GET("/stocks", func(c *gin.Context) {
		stocks := svc.ListStocks()
		c.JSON(http.StatusOK, stocks)
	})
}
