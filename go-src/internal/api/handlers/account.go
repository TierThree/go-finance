package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tierthree/go-finance/internal/db"
)

type account struct {
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
}

var accounts = []account{
	{Name: "Groceries", Balance: 1000.00},
	{Name: "Bills", Balance: 2000.00},
}

func ListAccounts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, accounts)
}

func CreateAccounts(c *gin.Context) {
	var ingestedAccount account

	if c.BindJSON(&ingestedAccount) == nil {
		if (ingestedAccount.Name == "") {
			c.String(http.StatusInternalServerError, "Please specify a name for the account")
		} else {
			err := db.CreateAccount(ingestedAccount.Name, ingestedAccount.Balance)
			if err != nil {
				c.String(http.StatusInternalServerError, "Failure")
			} else {
				c.String(http.StatusOK, "Success")
			}
		}
	} else {
		c.String(http.StatusInternalServerError, "Failure")
	}
}
