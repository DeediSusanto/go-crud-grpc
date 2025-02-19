package handler

import (
	"go-crud-grpc/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get all leads
func GetLeads(c *gin.Context, db *gorm.DB) {
	var leads []model.Lead
	if err := db.Find(&leads).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch leads"})
		return
	}
	c.JSON(http.StatusOK, leads)
}

// Get a single lead by ID
func GetLead(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var lead model.Lead

	if err := db.First(&lead, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
		return
	}
	c.JSON(http.StatusOK, lead)
}

// Create a new lead
func CreateLead(c *gin.Context, db *gorm.DB) {
	var lead model.Lead
	if err := c.ShouldBindJSON(&lead); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.Create(&lead).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create lead"})
		return
	}
	c.JSON(http.StatusCreated, lead)
}

// Update an existing lead
func UpdateLead(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var lead model.Lead

	if err := db.First(&lead, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
		return
	}

	if err := c.ShouldBindJSON(&lead); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.Save(&lead).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lead"})
		return
	}

	c.JSON(http.StatusOK, lead)
}

// Delete a lead
func DeleteLead(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var lead model.Lead

	if err := db.First(&lead, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
		return
	}

	if err := db.Delete(&lead).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete lead"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lead deleted successfully"})
}
