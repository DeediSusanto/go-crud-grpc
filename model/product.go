package model

import "gorm.io/gorm"

// Product represents the product table in the database
type Product struct {
	gorm.Model          // Automatically adds ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
