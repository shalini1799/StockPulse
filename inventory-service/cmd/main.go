package main

import (
	"github.com/gin-gonic/gin"
	handlers "inventory-service/internal/api"
	"inventory-service/internal/db"
	"github.com/joho/godotenv"
	"log"
	"inventory-service/internal/models"
)

func main() {

	// Load environment variables
	if err := godotenv.Load(); err != nil {
    log.Println("No .env file found inside container, using docker-compose environment variables.")
	}
	// Initialize the database connection
	db := db.InitDB()
	db.AutoMigrate(&models.Inventory{})

	inventoryHandler := handlers.InventoryHandler{DB: db}

	router := gin.Default()

	v1 := router.Group("/v1/inventory") 
	{
		v1.POST("/create", inventoryHandler.CreateInventory)
		v1.GET("/fetch-stock/:id", inventoryHandler.GetStock)
		v1.GET("/fetch-all-stock", inventoryHandler.GetAllStock)
		v1.PUT("/update-stock/:id", inventoryHandler.UpdateStock)
		v1.DELETE("/delete-stock/:id", inventoryHandler.DeleteStock)	

	}

}