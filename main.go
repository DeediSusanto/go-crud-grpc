package main

import (
	"fmt"
	"go-crud-grpc/config"
	"go-crud-grpc/handler"
	"go-crud-grpc/model"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Establish the connection to the database
	db, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Automatic migration for the User and Product tables
	err = db.AutoMigrate(&model.User{}, &model.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Connected to the database successfully!")

	// Create Gin router
	r := gin.Default()

	// Add routes for Users
	r.GET("/users", func(c *gin.Context) {
		handler.GetUsers(c, db) // Pass db to the handler
	})
	r.POST("/users", func(c *gin.Context) {
		handler.CreateUser(c, db) // Pass db to the handler
	})

	// Add routes for Products
	r.GET("/products", func(c *gin.Context) {
		handler.GetProducts(c, db) // Pass db to the handler
	})
	r.GET("/products/:id", func(c *gin.Context) {
		handler.GetProduct(c, db) // Pass db to the handler
	})
	r.POST("/products", func(c *gin.Context) {
		handler.CreateProduct(c, db) // Pass db to the handler
	})
	r.PUT("/products/:id", func(c *gin.Context) {
		handler.UpdateProduct(c, db) // Pass db to the handler
	})
	r.DELETE("/products/:id", func(c *gin.Context) {
		handler.DeleteProduct(c, db) // Pass db to the handler
	})

	// Run the server on port 8080
	r.Run(":8080")
}
