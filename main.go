package main

import (
	"go-crud-grpc/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Membuat router Gin
	r := gin.Default()

	// Menambahkan routing untuk API
	r.GET("/users", handler.GetUsers)
	r.POST("/users", handler.CreateUser)

	// Menjalankan server pada port 8080
	r.Run(":8080")
}
