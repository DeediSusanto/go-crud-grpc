package repository

import (
	"go-crud-grpc/config"
	"go-crud-grpc/model"
)

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]model.User, error) {
	var users []model.User
	db, err := config.ConnectToDatabase() // Connect to the database
	if err != nil {
		return nil, err
	}

	// Retrieve all users
	err = db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUsersByName retrieves users by name from the database
func GetUsersByName(name string) ([]model.User, error) {
	var users []model.User
	db, err := config.ConnectToDatabase() // Connect to the database
	if err != nil {
		return nil, err
	}

	// Filter by name using a LIKE query
	err = db.Where("name LIKE ?", "%"+name+"%").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
