package handlers

import (
	"net/conv"

	"github.com/gin-gonic/gin"
)

func ListBudgets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "true"})
}
