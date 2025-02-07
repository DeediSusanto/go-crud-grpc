package repository

import (
	"go-crud-grpc/config"
	"go-crud-grpc/model"
)

// GetUsersByName mendapatkan user berdasarkan nama
func GetUsersByName(name string) ([]model.User, error) {
	var users []model.User
	db, err := config.ConnectToDatabase()
	if err != nil {
		return nil, err
	}

	// Filter berdasarkan nama
	err = db.Where("name LIKE ?", "%"+name+"%").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
