package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tierthree/go-finance/internal/db"
)

type Account struct {
	Name    string  `json:"name" binding:"required"`
	Balance float32 `json:"balance"`
}

type AccountUpdate struct {
	Name       string  `json:"name" binding:"required"`
	NewName    string  `json:"new_name"`
	NewBalance float32 `json:"new_balance"`
}

func (au *AccountUpdate) InitDefaults() {
	au.Name = ""
	au.NewName = ""
	au.NewBalance = -1.0
}

func ListAccount(c *gin.Context) {
	var accounts = []Account{}
	accountDetails, _ := db.ReadAccounts()

	for accountName, accountBalance := range accountDetails {
		accounts = append(accounts, Account{Name: accountName, Balance: accountBalance})
	}

	c.IndentedJSON(http.StatusOK, accounts)
}

func CreateAccount(c *gin.Context) {
	var ingestedAccount Account

	if c.BindJSON(&ingestedAccount) == nil {
		err := db.CreateAccount(ingestedAccount.Name, ingestedAccount.Balance)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failure")
		} else {
			c.String(http.StatusOK, "Success")
		}
	} else {
		c.String(http.StatusInternalServerError, "Failure")
	}
}

func DeleteAccount(c *gin.Context) {
	var ingestedAccount Account

	if c.BindJSON(&ingestedAccount) == nil {
		err := db.DeleteAccount(ingestedAccount.Name)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failure")
		} else {
			c.String(http.StatusOK, "Success")
		}
	} else {
		c.String(http.StatusInternalServerError, "Failure")
	}
}

func UpdateAccount(c *gin.Context) {
	var ingestedAccountUpdate AccountUpdate
	ingestedAccountUpdate.InitDefaults()

	if c.BindJSON(&ingestedAccountUpdate) == nil {
		if ingestedAccountUpdate.NewBalance != -1.0 {
			err := db.UpdateAccountBalance(
				ingestedAccountUpdate.Name,
				ingestedAccountUpdate.NewBalance,
			)
			ValidateError(c, err)
		}

		if ingestedAccountUpdate.NewName != "" {
			err := db.UpdateAccountName(
				ingestedAccountUpdate.Name,
				ingestedAccountUpdate.NewName,
			)
			ValidateError(c, err)
		}
	} else {
		c.String(http.StatusInternalServerError, "Failure")
	}
}

func ValidateError(c *gin.Context, err error) {
	if err != nil {
		c.String(http.StatusInternalServerError, "Failure")
	} else {
		c.String(http.StatusOK, "Success")
	}
}
