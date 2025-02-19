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

	// Automatic migration for the User, Product, and Lead tables
	err = db.AutoMigrate(&model.User{}, &model.Product{}, &model.Lead{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Connected to the database successfully!")

	// Create Gin router
	r := gin.Default()

	// Add routes for Users
	r.GET("/users", func(c *gin.Context) {
		handler.GetUsers(c, db)
	})
	r.POST("/users", func(c *gin.Context) {
		handler.CreateUser(c, db)
	})

	// Add routes for Products
	r.GET("/products", func(c *gin.Context) {
		handler.GetProducts(c, db)
	})
	r.GET("/products/:id", func(c *gin.Context) {
		handler.GetProduct(c, db)
	})
	r.POST("/products", func(c *gin.Context) {
		handler.CreateProduct(c, db)
	})
	r.PUT("/products/:id", func(c *gin.Context) {
		handler.UpdateProduct(c, db)
	})
	r.DELETE("/products/:id", func(c *gin.Context) {
		handler.DeleteProduct(c, db)
	})

	// Add routes for Leads
	r.GET("/leads", func(c *gin.Context) {
		handler.GetLeads(c, db)
	})
	r.GET("/leads/:id", func(c *gin.Context) {
		handler.GetLead(c, db)
	})
	r.POST("/leads", func(c *gin.Context) {
		handler.CreateLead(c, db)
	})
	r.PUT("/leads/:id", func(c *gin.Context) {
		handler.UpdateLead(c, db)
	})
	r.DELETE("/leads/:id", func(c *gin.Context) {
		handler.DeleteLead(c, db)
	})

	// Run the server on port 8080
	r.Run(":8080")
}
