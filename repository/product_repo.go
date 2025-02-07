package repository

import (
	"database/sql"       // For MySQL
	"go-crud-grpc/model" // Make sure this import is correct based on your project structure
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Example function to fetch a product by ID
func (repo *ProductRepository) GetProduct(id int32) (*model.Product, error) {
	query := "SELECT id, name, description, price FROM products WHERE id = ?"
	row := repo.db.QueryRow(query, id)

	var product model.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
