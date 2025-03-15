package service

import (
	"fmt"
	"go-crud-grpc/model"
	"go-crud-grpc/repository"
	"go-crud-grpc/utils"
)

type TransactionService struct {
	repo        *repository.TransactionRepository
	userRepo    *repository.UserRepository
	productRepo *repository.ProductRepository
}

func NewTransactionService(
	repo *repository.TransactionRepository,
	userRepo *repository.UserRepository,
	productRepo *repository.ProductRepository,
) *TransactionService {
	return &TransactionService{
		repo:        repo,
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}

// ✅ CreateTransaction dengan validasi user & produk
func (s *TransactionService) CreateTransaction(userID uint, productID uint, amount float64, cardNumber string) (*model.Transaction, error) {
	// ✅ Validasi jumlah transaksi
	if amount <= 0 {
		return nil, fmt.Errorf("invalid transaction amount: must be greater than 0")
	}

	// ✅ Cek apakah user_id ada
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user with ID %d not found", userID)
	}

	// ✅ Cek apakah product_id ada
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		return nil, fmt.Errorf("product with ID %d not found", productID)
	}

	// ✅ Tokenisasi kartu kredit
	token, err := utils.TokenizeCard(cardNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to tokenize card: %v", err)
	}

	// ✅ Buat transaksi baru
	transaction := &model.Transaction{
		UserID:    user.ID,
		ProductID: uint(product.ID), // ✅ Konversi dari int32 ke uint
		Amount:    amount,
		CardToken: token,
		Status:    "Pending",
	}

	// ✅ Simpan transaksi
	if err := s.repo.Create(transaction); err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}

	return transaction, nil
}

// ✅ Get semua transaksi
func (s *TransactionService) GetAllTransactions() ([]model.Transaction, error) {
	transactions, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transactions: %v", err)
	}
	return transactions, nil
}

// ✅ Get transaksi berdasarkan ID
func (s *TransactionService) GetTransactionByID(id uint) (*model.Transaction, error) {
	transaction, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("transaction with ID %d not found", id)
	}
	return transaction, nil
}
