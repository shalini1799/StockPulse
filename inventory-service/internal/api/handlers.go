package api

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"inventory-service/internal/models"
	"gorm.io/gorm"
	"time"
	"strconv"
)

type InventoryHandler struct {
	DB *gorm.DB
}

// CreateInventory handles the creation of a new inventory item
func (handler *InventoryHandler) CreateInventory(c *gin.Context) {
	tx := handler.DB.Begin()

	inventory := models.Inventory{}
	if err := c.ShouldBindJSON(&inventory); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tx.Create(&inventory).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create inventory"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Inventory created successfully", "inventory": inventory})
}


func (handler *InventoryHandler) GetStock(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Inventory
	if err := handler.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product details fetched successfully", "inventory": product})
}

func (handler *InventoryHandler) GetAllStock(c *gin.Context) {
	var products []models.Inventory

	if err := handler.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All products fetched successfully", "inventory": products})
}

func (handler *InventoryHandler) UpdateStock(c *gin.Context) {
	tx := handler.DB.Begin()

	var updated models.Inventory
	if err := c.ShouldBindJSON(&updated); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated.UpdatedAt = time.Now()

	if err := tx.Model(&models.Inventory{}).Where("id = ?", updated.Id).Updates(updated).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update product"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Product details updated successfully", "inventory": updated})
}

func (handler *InventoryHandler) DeleteStock(c *gin.Context) {
	tx := handler.DB.Begin()

	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Inventory
	if err := tx.First(&product, productID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := tx.Delete(&product).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete product"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"message":  "Product deleted successfully",
		"inventory": product,
	})
}