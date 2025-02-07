# go-crud-grpc
 Best Pratice & Server Pattern Sharing in Golang with Study Case


 Best practices menggunakan service pattern di Golang dan gRPC API untuk CRUD Produk. Di sini, kita akan menggunakan struktur folder yang lebih modular dan bersih, serta mengikuti pola service untuk pemisahan tanggung jawab dan testabilitas yang lebih baik.

Berikut adalah langkah-langkah yang akan kita lakukan:

Struktur Folder
Sebelum masuk ke implementasi, mari kita tentukan struktur folder terlebih dahulu:
/go-crud-grpc
├── /cmd
│   └── main.go               # Entry point untuk aplikasi
├── /config                   # File untuk mengonfigurasi database dan service
│   └── database.go
├── /model                    # Model data yang digunakan oleh GORM
│   └── product.go
├── /repository               # CRUD Repository untuk berinteraksi dengan DB
│   └── product_repo.go
├── /service                  # Layanan untuk produk yang memisahkan logika bisnis
│   └── product_service.go
├── /handler                  # Handler HTTP untuk CRUD API (REST)
│   └── product_handler.go
├── /pb                       # File gRPC definitions (Protocol Buffers)
│   └── product.proto
├── /server                   # Server gRPC
│   └── grpc_server.go
├── go.mod
├── go.sum
└── README.md


1. Menyiapkan File product.proto
Langkah pertama adalah membuat gRPC service definition menggunakan file Protocol Buffers (product.proto). File ini akan mendefinisikan struktur data dan service gRPC kita.

/pb/product.proto:

proto
Copy
syntax = "proto3";

package product;

service ProductService {
    rpc CreateProduct (CreateProductRequest) returns (ProductResponse);
    rpc GetProduct (GetProductRequest) returns (ProductResponse);
    rpc GetAllProducts (GetAllProductsRequest) returns (ProductListResponse);
    rpc UpdateProduct (UpdateProductRequest) returns (ProductResponse);
    rpc DeleteProduct (DeleteProductRequest) returns (ProductResponse);
}

message Product {
    int32 id = 1;
    string name = 2;
    string description = 3;
    float price = 4;
}

message CreateProductRequest {
    string name = 1;
    string description = 2;
    float price = 3;
}

message GetProductRequest {
    int32 id = 1;
}

message UpdateProductRequest {
    int32 id = 1;
    string name = 2;
    string description = 3;
    float price = 4;
}

message DeleteProductRequest {
    int32 id = 1;
}

message ProductResponse {
    bool success = 1;
    string message = 2;
    Product product = 3;
}

message GetAllProductsRequest {}

message ProductListResponse {
    repeated Product products = 1;
}
Penjelasan:
ProductService adalah service gRPC yang mendefinisikan lima metode: CreateProduct, GetProduct, GetAllProducts, UpdateProduct, dan DeleteProduct.
Product adalah model data untuk produk yang memiliki id, name, description, dan price.
Request dan Response mendefinisikan data yang akan dikirim dan diterima dalam setiap metode.
2. Generate gRPC Code
Setelah file product.proto selesai, kita perlu menghasilkan kode Go dari file .proto untuk gRPC. Kamu perlu menginstal protoc dan plugin Go untuk gRPC jika belum terpasang.

Install protoc (Protocol Buffers Compiler):

Download dan install protoc dari Protoc Releases.
Install plugin Go untuk gRPC:

Jalankan perintah ini di terminal untuk menginstal plugin Go:

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
Generate kode gRPC: Setelah itu, jalankan perintah berikut untuk menghasilkan file Go dari product.proto:

protoc --go_out=. --go-grpc_out=. pb/product.proto
Perintah ini akan menghasilkan dua file Go di folder /pb:

product.pb.go (untuk pesan dan definisi struktur data)
product_grpc.pb.go (untuk implementasi gRPC)
3. Menyiapkan Model Product
Di folder model, buat model Product yang akan berinteraksi dengan database menggunakan GORM.

/model/product.go:

package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
4. Repository untuk Akses Database
Di folder repository, buat file untuk melakukan operasi CRUD pada database menggunakan GORM.

/repository/product_repo.go:


package repository

import (
	"go-crud-grpc/config"
	"go-crud-grpc/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) (*model.Product, error)
	GetById(id int32) (*model.Product, error)
	GetAll() ([]model.Product, error)
	Update(product *model.Product) (*model.Product, error)
	Delete(id int32) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository() ProductRepository {
	db, _ := config.ConnectToDatabase()
	return &productRepo{db}
}

func (r *productRepo) Create(product *model.Product) (*model.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *productRepo) GetById(id int32) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *productRepo) GetAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepo) Update(product *model.Product) (*model.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *productRepo) Delete(id int32) error {
	return r.db.Delete(&model.Product{}, id).Error
}
5. Service untuk Logika Bisnis
Di folder service, buat service yang akan memproses logika bisnis untuk CRUD produk.

/service/product_service.go:


package service

import (
	"go-crud-grpc/model"
	"go-crud-grpc/repository"
	"golang.org/x/net/context"
	"go-crud-grpc/pb"
)

type ProductService struct {
	productRepo repository.ProductRepository
	pb.UnimplementedProductServiceServer
}

func NewProductService() *ProductService {
	return &ProductService{
		productRepo: repository.NewProductRepository(),
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	product := model.Product{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
	}
	createdProduct, err := s.productRepo.Create(&product)
	if err != nil {
		return &pb.ProductResponse{
			Success: false,
			Message: "Failed to create product",
		}, err
	}

	return &pb.ProductResponse{
		Success: true,
		Message: "Product created successfully",
		Product: &pb.Product{
			Id:          int32(createdProduct.ID),
			Name:        createdProduct.Name,
			Description: createdProduct.Description,
			Price:       float32(createdProduct.Price),
		},
	}, nil
}

// Implementasi metode lainnya (GetProduct, GetAllProducts, UpdateProduct, DeleteProduct) bisa ditambahkan di sini.
6. Handler untuk HTTP API (REST)
Di folder handler, buat file handler untuk menangani API REST yang memungkinkan pengguna berinteraksi dengan CRUD produk via HTTP.

/handler/product_handler.go:


package handler

import (
	"go-crud-grpc/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req service.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result, err := h.service.CreateProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
7. Server gRPC
Akhirnya, kita akan membuat server gRPC yang akan menjalankan service ProductService.

/server/grpc_server.go:


package server

import (
	"go-crud-grpc/pb"
	"go-crud-grpc/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, service.NewProductService())

	log.Println("gRPC server listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
8. Main Function
Terakhir, kita akan menambahkan fungsi main.go sebagai entry point untuk aplikasi.

/cmd/main.go:


package main

import (
	"go-crud-grpc/server"
)

func main() {
	go server.StartGrpcServer()
	// Tambahkan juga server HTTP atau fitur lainnya jika diperlukan.
}
9. Testing dengan Postman
Untuk menguji API REST, kamu bisa menggunakan Postman dengan menambahkan endpoint seperti berikut:

bash
Copy
POST http://localhost:8080/products
Untuk menguji gRPC, kamu bisa menggunakan Postman atau tools gRPC lainnya seperti Insomnia.

