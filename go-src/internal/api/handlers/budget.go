package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tierthree/go-finance/internal/db"
)

type Budget struct {
	Name    string  `json:"name" binding:"required"`
	Amount float32 `json:"amount"`
}

type BudgetUpdate struct {
	Name       string  `json:"name" binding:"required"`
	NewName    string  `json:"new_name"`
	NewAmount float32 `json:"new_amount"`
}

func (au *BudgetUpdate) InitDefaults() {
	au.Name = ""
	au.NewName = ""
	au.NewAmount = -1.0
}

func ListBudget(c *gin.Context) {
	var budgets = []Budget{}
	budgetDetails, _ := db.ReadBudgets()

	for budgetName, budgetAmount := range budgetDetails {
		budgets = append(budgets, Budget{Name: budgetName, Amount: budgetAmount})
	}

	c.IndentedJSON(http.StatusOK, budgets)
}

func CreateBudget(c *gin.Context) {
	var ingestedBudget Budget

	if c.BindJSON(&ingestedBudget) == nil {
		err := db.CreateBudget(ingestedBudget.Name, ingestedBudget.Amount)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failure")
		} else {
			c.String(http.StatusOK, "Success")
		}
	} else {
		c.String(http.StatusInternalServerError, "Failure to BindJSON")
	}
}

func DeleteBudget(c *gin.Context) {
	var ingestedBudget Budget

	if c.BindJSON(&ingestedBudget) == nil {
		err := db.DeleteBudget(ingestedBudget.Name)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failure")
		} else {
			c.String(http.StatusOK, "Success")
		}
	} else {
		c.String(http.StatusInternalServerError, "Failure to BindJSON")
	}
}

func UpdateBudget(c *gin.Context) {
	var ingestedBudgetUpdate BudgetUpdate
	ingestedBudgetUpdate.InitDefaults()

	if c.BindJSON(&ingestedBudgetUpdate) == nil {
		if ingestedBudgetUpdate.NewAmount != -1.0 {
			err := db.UpdateBudgetAmount(
				ingestedBudgetUpdate.Name,
				ingestedBudgetUpdate.NewAmount,
			)
			ValidateError(c, err)
		}

		if ingestedBudgetUpdate.NewName != "" {
			err := db.UpdateBudgetName(
				ingestedBudgetUpdate.Name,
				ingestedBudgetUpdate.NewName,
			)
			ValidateError(c, err)
		}
	} else {
		c.String(http.StatusInternalServerError, "Failure to BindJSON")
	}
}
