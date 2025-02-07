package handler

import (
	"go-crud-grpc/model"
	"go-crud-grpc/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUsers handler to retrieve all users
func GetUsers(c *gin.Context, db *gorm.DB) {
	// Get the query parameter "name" if available
	name := c.DefaultQuery("name", "") // Default empty if no query parameter "name"

	var users []model.User
	var err error

	// If there's a "name" query parameter, fetch users by name
	if name != "" {
		users, err = repository.GetUsersByName(name)
	} else {
		// Otherwise, fetch all users
		users, err = repository.GetAllUsers()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	// If no users found, return 404 or 204
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	// If users are found, return them in JSON format
	c.JSON(http.StatusOK, users)
}

// CreateUser handler to create a user
func CreateUser(c *gin.Context, db *gorm.DB) {
	var user model.User

	// Bind the incoming JSON to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Save the user in the database
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
