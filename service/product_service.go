package service

import (
	"go-crud-grpc/model"
	"go-crud-grpc/pb"
	"go-crud-grpc/repository"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductService struct {
	productRepo *repository.ProductRepository
	pb.UnimplementedProductServiceServer
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{
		productRepo: repository.NewProductRepository(db),
	}
}

// ✅ Create Product
func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	// Cek apakah nama produk kosong
	if req.GetName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Product name cannot be empty")
	}

	product := model.Product{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       float64(req.GetPrice()), // ✅ Konversi float32 ke float64
	}

	if err := s.productRepo.Create(&product); err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Success: true,
		Message: "Product created successfully",
		Product: &pb.Product{
			Id:          int32(product.ID), // ✅ Pastikan ID tetap int32
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price), // ✅ Konversi float64 ke float32
		},
	}, nil

}
