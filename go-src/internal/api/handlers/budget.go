package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type budget struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
}

var budgets = []budget{
	{Name: "Groceries", Amount: 1000.00},
	{Name: "Bills", Amount: 2000.00},
}

func ListBudgets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, budgets)
}

func CreateBudgets(c *gin.Context) {
	var ingestedBudget budget

	if c.BindJSON(&ingestedBudget) == nil {
		budgets = append(budgets, ingestedBudget)
		c.String(http.StatusOK, "Success")
	} else {
		c.String(http.StatusInternalServerError, "Failure")
	}
}
