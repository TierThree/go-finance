package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tierthree/go-finance/internal/api/routes"
	"github.com/tierthree/go-finance/internal/db"
)

func main() {
	db.ReadAccountBalance("test")

	router := gin.Default()

	routerGroup := router.Group("/api/v1")
	routes.BudgetRouter(routerGroup)
	routes.TransactionRouter(routerGroup)
	routes.AccountRouter(routerGroup)

	router.Run("localhost:8080")
}
