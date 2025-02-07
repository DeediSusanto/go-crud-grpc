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
