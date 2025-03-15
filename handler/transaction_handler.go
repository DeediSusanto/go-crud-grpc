package handler

import (
	"go-crud-grpc/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	// ✅ Gunakan binding JSON agar lebih aman & fleksibel
	var request struct {
		UserID     uint    `json:"user_id"`
		ProductID  uint    `json:"product_id"`
		Amount     float64 `json:"amount"`
		CardNumber string  `json:"card_number"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// ✅ Debugging: Print nilai amount (jika diperlukan)
	// fmt.Println("Amount received:", request.Amount)

	transaction, err := h.service.CreateTransaction(request.UserID, request.ProductID, request.Amount, request.CardNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := h.service.GetTransactionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
