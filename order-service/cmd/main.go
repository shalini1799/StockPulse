// order-service/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"order-service/internal/db"
	"log"
	"github.com/joho/godotenv"
	handlers "order-service/internal/api" 
	request_models "order-service/internal/models/request_models"
)

func main() {
	if err := godotenv.Load(); err != nil {
    log.Println("No .env file found inside container, using docker-compose environment variables.")
	}

	db := db.InitDB()
	//create the orders table if it doesn't exist
	db.AutoMigrate(&request_models.Order{})

	orderService := &handlers.OrderService{DB: db}

	router := gin.Default()
	router.POST("/order", orderService.CreateOrder)

	router.Run(":8080")
}