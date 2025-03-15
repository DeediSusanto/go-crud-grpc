package repository

import (
	"go-crud-grpc/model"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// ✅ Create Transaction
func (r *TransactionRepository) Create(transaction *model.Transaction) error {
	return r.DB.Create(transaction).Error
}

// ✅ Get Transaction by ID
func (r *TransactionRepository) GetByID(id uint) (*model.Transaction, error) {
	var transaction model.Transaction
	err := r.DB.First(&transaction, id).Error
	return &transaction, err
}

// ✅ Get All Transactions
func (r *TransactionRepository) GetAll() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.DB.Find(&transactions).Error
	return transactions, err
}
