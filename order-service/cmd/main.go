package main

import (
	"github.com/gin-gonic/gin"
	"stockpulse/internal/api"
)

func main() {
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "StockPulse backend is healthy and running!"})
	})

	// Stock routes
	api.RegisterRoutes(router)

	router.Run(":8080")
}