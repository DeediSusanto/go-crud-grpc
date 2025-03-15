package main

import (
	"fmt"
	"go-crud-grpc/config"
	"go-crud-grpc/handler"
	"go-crud-grpc/model"
	"go-crud-grpc/repository"
	"go-crud-grpc/service"
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

	// ✅ Create repositories
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	leadRepo := repository.NewLeadRepository() // ✅ Tanpa db
	transactionRepo := repository.NewTransactionRepository(db)

	// ✅ Create services
	transactionService := service.NewTransactionService(transactionRepo, userRepo, productRepo) // ✅ Perbaikan

	// ✅ Create handlers
	userHandler := handler.NewUserHandler(userRepo)
	productHandler := handler.NewProductHandler(productRepo)
	leadHandler := handler.NewLeadHandler(leadRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// ✅ Routes for Users
	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)

	// ✅ Routes for Products
	r.GET("/products", productHandler.GetProducts)
	r.GET("/products/:id", productHandler.GetProduct)
	r.POST("/products", productHandler.CreateProduct)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	// ✅ Routes for Leads
	r.GET("/leads", leadHandler.GetLeads)
	r.GET("/leads/:id", leadHandler.GetLead)
	r.POST("/leads", leadHandler.CreateLead)
	r.PUT("/leads/:id", leadHandler.UpdateLead)
	r.DELETE("/leads/:id", leadHandler.DeleteLead)

	// ✅ Routes for Transactions
	r.POST("/transactions", transactionHandler.CreateTransaction)
	r.GET("/transactions", transactionHandler.GetTransactions)
	r.GET("/transactions/:id", transactionHandler.GetTransactionByID)

	// Run the server on port 8080
	r.Run(":8080")
}
