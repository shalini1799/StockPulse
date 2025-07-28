// order-service/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"order-service/internal/db"
	"log"
	"github.com/joho/godotenv"
	handlers "order-service/internal/api" 
	"order-service/internal/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
    log.Println("No .env file found inside container, using docker-compose environment variables.")
	}

	db.InitDB()
	db.DB.AutoMigrate(&models.Order{})

	router := gin.Default()
	router.POST("/order", handlers.CreateOrder)
	router.GET("/order/:id", handlers.GetOrder)

	router.Run(":8080")
}