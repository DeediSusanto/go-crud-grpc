package model

import "gorm.io/gorm"

// Model User yang akan dipetakan ke dalam tabel users di database
type User struct {
	gorm.Model        // GORM akan menambahkan ID, CreatedAt, UpdatedAt, dan DeletedAt
	Name       string `json:"name"`
	Email      string `json:"email"`
}
