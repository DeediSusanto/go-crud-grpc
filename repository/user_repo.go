package repository

import (
	"go-crud-grpc/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// ✅ Get User by ID
func (repo *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := repo.DB.First(&user, id).Error
	return &user, err
}

// ✅ Get All Users
func (repo *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := repo.DB.Find(&users).Error
	return users, err
}

// ✅ Get Users by Name
func (repo *UserRepository) GetByName(name string) ([]model.User, error) {
	var users []model.User
	err := repo.DB.Where("name LIKE ?", "%"+name+"%").Find(&users).Error
	return users, err
}
