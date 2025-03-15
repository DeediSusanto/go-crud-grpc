package routes

import (
	"go-pci-transactions/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, transactionHandler *handler.TransactionHandler) {
	router.POST("/transactions", transactionHandler.CreateTransaction)
}
