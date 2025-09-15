package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tierthree/go-finance/internal/db"
)

type Transaction struct {
	Name    string    `json:"name" binding:"required"`
	Budget  string    `json:"budget" binding:"required"`
	Account string    `json:"account" binding:"required"`
	Amount  float32   `json:"amount" binding:"required"`
	Date    string    `json:"date" binding:"required"`
}

//var transactions = []Transaction{
//	{Name: "Electric Bill", Budget: "Bills", Account: "Test", Amount: 100.00, Date: time.Date(2025, time.September, 8, 0, 0, 0, 0, time.UTC)},
//}

func ListTransaction(c *gin.Context) {
	var transactions = []Transaction{}
	transactionDetails, _ := db.ReadTransactions()

	for _, transactionMap := range transactionDetails {
		amount, _ := strconv.ParseFloat(transactionMap["Amount"], 32)
		transaction := Transaction{
			Name: transactionMap["Name"],
			Budget: transactionMap["Budget"],
			Account: transactionMap["Account"],
			Amount: float32(amount),
			Date: transactionMap["Date"],
		}
		transactions = append(transactions, transaction)
	}

	c.IndentedJSON(http.StatusOK, transactions)
}

func CreateTransaction(c *gin.Context) {
	var ingestedTransaction Transaction 

	if c.BindJSON(&ingestedTransaction) == nil {
		balance, err := db.ReadAccountBalance(ingestedTransaction.Account)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failure to read account")
		} else {
			err = db.CreateTransaction(
				ingestedTransaction.Name, 
				ingestedTransaction.Budget,
				ingestedTransaction.Account,
				ingestedTransaction.Amount,
				ingestedTransaction.Date,
			)
			if err != nil {
				c.String(http.StatusInternalServerError, "Failure")
			} else {
				db.UpdateAccountBalance(
					ingestedTransaction.Account,
					balance - ingestedTransaction.Amount,
				)
				c.String(http.StatusOK, "Success")
			}

			fmt.Println(ingestedTransaction.Date)
		}
	} else {
		c.String(http.StatusInternalServerError, "Failure")
	}
}
