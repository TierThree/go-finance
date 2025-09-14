package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tierthree/go-finance/internal/api/handlers"
)

func BudgetRouter(router *gin.RouterGroup) {
	budget := router.Group("/budget")

	budget.GET("/list", handlers.ListBudget)
	budget.POST("/create", handlers.CreateBudget)
	budget.POST("/delete", handlers.DeleteBudget)
	budget.POST("/update", handlers.UpdateBudget)
}
