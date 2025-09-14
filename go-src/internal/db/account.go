package db

import (
	"fmt"
)

var tableDetails = table{
	name:   "accounts",
	fields: "name,balance",
}

func CreateAccount(name string, balance float32) error {
	tableParam := fmt.Sprintf("%s (%s)", tableDetails.name, tableDetails.fields)
	values := fmt.Sprintf("'%s', %0.2f", name, balance)

	fmt.Printf("%s %s\n", tableParam, values)

	err := Create(tableParam, values)

	return err
}

func DeleteAccount(name string) error {
	condition := fmt.Sprintf("name = '%s'", name)

	err := Delete(tableDetails.name, condition)

	return err
}

func UpdateAccountName(name string, newName string) error {
	updatedValue := fmt.Sprintf("name = '%s'", newName)

	err := Update(tableDetails.name, updatedValue, name)

	return err
}

func UpdateAccountBalance(name string, newBalance float32) error {
	updatedValue := fmt.Sprintf("balance = %0.2f", newBalance)

	err := Update(tableDetails.name, updatedValue, name)

	return err
}

func ReadAccountBalance(name string) (float32, error) {
	filter := fmt.Sprintf("name = '%s'", name)

	row, err := ReadSingleRow(tableDetails.name, "balance", filter)
	if err != nil {
		return 0.0, fmt.Errorf("Error reading from DB: %w", err)
	}

	var balance float32
	row.Scan(&balance)

	fmt.Printf("Balance for account %s is: %0.2f\n", name, balance)

	return balance, nil
}

func ReadAccounts() (map[string]float32, error) {
	accountDetails := make(map[string]float32)

	rows, err := ReadMultipleRows(tableDetails.name, tableDetails.fields)
	if err != nil {
		return nil, fmt.Errorf("Error querying table: %w", err)
	}

	for rows.Next() {
		var accountName string
		var accountBalance float32

		rows.Scan(&accountName, &accountBalance)

		accountDetails[accountName] = accountBalance
	}

	return accountDetails, nil
}
