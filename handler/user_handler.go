package handler

import (
	"net/http"

	"go-crud-grpc/model"

	"github.com/gin-gonic/gin"
)

// GetUsers handler untuk mengambil daftar user
func GetUsers(c *gin.Context) {
	users := []model.User{
		{ID: 1, Name: "Dedi", Email: "dedi@example.com"},
		{ID: 2, Name: "Budi", Email: "budi@example.com"},
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser handler untuk menambahkan user baru
func CreateUser(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Di sini kamu bisa menambahkan logic untuk menyimpan user di database

	c.JSON(http.StatusCreated, newUser)
}
