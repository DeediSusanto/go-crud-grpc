package handler

import (
	"go-crud-grpc/model"
	"go-crud-grpc/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers handler untuk mengambil semua user
func GetUsers(c *gin.Context) {
	// Mendapatkan query parameter "name" jika ada
	name := c.DefaultQuery("name", "") // Default kosong jika tidak ada query parameter "name"

	// Jika ada query parameter name, kita filter berdasarkan nama
	var users []model.User
	var err error

	if name != "" {
		// Jika ada nama, ambil user berdasarkan nama
		users, err = repository.GetUsersByName(name)
	} else {
		// Jika tidak ada nama, ambil semua user
		users, err = repository.GetAllUsers()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	// Jika data kosong, kembalikan pesan dengan status 404 atau 204
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	// Jika ada data, kembalikan hasil dalam format JSON
	c.JSON(http.StatusOK, users)
}
