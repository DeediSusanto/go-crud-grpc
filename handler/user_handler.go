package handler

import (
	"go-crud-grpc/model"
	"go-crud-grpc/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

// ✅ Get Users
func (h *UserHandler) GetUsers(c *gin.Context) {
	name := c.DefaultQuery("name", "")

	var users []model.User
	var err error

	if name != "" {
		users, err = h.repo.GetByName(name) // 🔄 Sesuai dengan method di UserRepository
	} else {
		users, err = h.repo.GetAll() // 🔄 Sesuai dengan method di UserRepository
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// ✅ Create User
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.repo.DB.Create(&user).Error; err != nil { // 🔄 Sesuai dengan penggunaan GORM
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
