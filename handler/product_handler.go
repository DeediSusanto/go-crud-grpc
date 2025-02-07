package handler

import (
	"go-crud-grpc/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProducts handler to retrieve all products
func GetProducts(c *gin.Context, db *gorm.DB) {
	var products []model.Product
	err := db.Find(&products).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProduct handler to retrieve a product by ID
func GetProduct(c *gin.Context, db *gorm.DB) {
	// Get the product ID from the URL
	id := c.Param("id")

	// Fetch the product by ID
	var product model.Product
	err := db.First(&product, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct handler to create a product
func CreateProduct(c *gin.Context, db *gorm.DB) {
	// Bind the incoming JSON to the product struct
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Save the product in the database
	if err := db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// UpdateProduct handler to update a product by ID
func UpdateProduct(c *gin.Context, db *gorm.DB) {
	// Get the product ID from the URL
	id := c.Param("id")

	// Fetch the product by ID
	var product model.Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Bind the incoming JSON to the product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Save the updated product in the database
	if err := db.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct handler to delete a product by ID
func DeleteProduct(c *gin.Context, db *gorm.DB) {
	// Get the product ID from the URL
	id := c.Param("id")

	// Fetch the product by ID
	var product model.Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Delete the product from the database
	if err := db.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
