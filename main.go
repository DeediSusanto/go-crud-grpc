package main

import (
	"log"

	"go-crud-grpc/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Membuat instance Gin router
	r := gin.Default()

	// Definisikan route untuk CRUD User
	r.GET("/users", handler.GetUsers)
	r.POST("/users", handler.CreateUser)

	// Menjalankan server di port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
