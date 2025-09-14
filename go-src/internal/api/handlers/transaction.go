package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Name   string
	Budget string
	Amount float32
	Date   string
}

var transactions = []Transaction{
	{Name: "Electric Bill", Budget: "Bills", Amount: 100.00, Date: "09/08/2025"},
}

func ListTransactions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, transactions)
}
