package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tierthree/go-finance/internal/api/handlers"
)

func TransactionRouter(router *gin.RouterGroup) {
	transaction := router.Group("/transaction")

	transaction.GET("/list", handlers.ListTransactions)
}
