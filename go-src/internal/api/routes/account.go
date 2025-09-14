package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tierthree/go-finance/internal/api/handlers"
)

func AccountRouter(router *gin.RouterGroup) {
	account := router.Group("/account")

	account.GET("/list", handlers.ListAccount)
	account.POST("/create", handlers.CreateAccount)
	account.POST("/delete", handlers.DeleteAccount)
	account.POST("/update", handlers.UpdateAccount)
}
