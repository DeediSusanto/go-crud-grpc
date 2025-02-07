package service

import (
	"go-crud-grpc/model"
	"go-crud-grpc/pb"
	"go-crud-grpc/repository"

	"golang.org/x/net/context"
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
