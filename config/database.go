package config

import (
	"fmt"
	"go-crud-grpc/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectToDatabase untuk koneksi ke MySQL dan migrasi tabel
func ConnectToDatabase() (*gorm.DB, error) {
	// String koneksi ke database MySQL (sesuaikan jika berbeda)
	dsn := "root:@tcp(localhost)/go_crud_db?charset=utf8mb4&parseTime=True&loc=Local"
	// Membuka koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrasi otomatis untuk model User
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}

	return db, nil
}
