package db

import (
	"fmt"
)

var tableBudgetDetails = table{
	name:   "budgets",
	fields: "name,amount",
}

func CreateBudget(name string, amount float32) error {
	tableParam := fmt.Sprintf("%s (%s)", tableBudgetDetails.name, tableBudgetDetails.fields)
	values := fmt.Sprintf("'%s', %0.2f", name, amount)

	fmt.Printf("%s %s\n", tableParam, values)

	err := Create(tableParam, values)

	return err
}

func DeleteBudget(name string) error {
	condition := fmt.Sprintf("name = '%s'", name)

	err := Delete(tableBudgetDetails.name, condition)

	return err
}

func UpdateBudgetName(name string, newName string) error {
	updatedValue := fmt.Sprintf("name = '%s'", newName)

	err := Update(tableBudgetDetails.name, updatedValue, name)

	return err
}

func UpdateBudgetAmount(name string, newAmount float32) error {
	updatedValue := fmt.Sprintf("amount = %0.2f", newAmount)

	err := Update(tableBudgetDetails.name, updatedValue, name)

	return err
}

func ReadBudgetAmount(name string) (float32, error) {
	filter := fmt.Sprintf("name = '%s'", name)

	row, err := ReadSingleRow(tableBudgetDetails.name, "amount", filter)
	if err != nil {
		return 0.0, fmt.Errorf("Error reading from DB: %w", err)
	}

	var amount float32
	row.Scan(&amount)

	return amount, nil
}

func ReadBudgets() (map[string]float32, error) {
	budgetDetails := make(map[string]float32)

	rows, err := ReadMultipleRows(tableBudgetDetails.name, tableBudgetDetails.fields)
	if err != nil {
		return nil, fmt.Errorf("Error querying table: %w", err)
	}

	for rows.Next() {
		var budgetName string
		var budgetAmount float32

		rows.Scan(&budgetName, &budgetAmount)

		budgetDetails[budgetName] = budgetAmount
	}

	return budgetDetails, nil
}
