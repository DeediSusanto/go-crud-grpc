package repository

import (
	"go-crud-grpc/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// ✅ Get All Products
func (r *ProductRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Find(&products).Error
	return products, err
}

// ✅ Get Product by ID
func (r *ProductRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.DB.First(&product, id).Error
	return &product, err
}

// ✅ Create Product
func (r *ProductRepository) Create(product *model.Product) error {
	return r.DB.Create(product).Error
}

// ✅ Update Product
func (r *ProductRepository) Update(product *model.Product) error {
	return r.DB.Save(product).Error
}

// ✅ Delete Product
func (r *ProductRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Product{}, id).Error
}
