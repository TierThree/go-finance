package db

import (
	"fmt"
	"time"
)

var tableTransactionDetails = table{
	name:   "transactions",
	fields: "name,amount,budget,account,date",
}

func CreateTransaction(name string, budget string, account string, amount float32, date string) error {
	tableParam := fmt.Sprintf("%s (%s)", tableTransactionDetails.name, tableTransactionDetails.fields)

	values := fmt.Sprintf("'%s', %0.2f, '%s', '%s', '%s'", name, amount, budget, account, date)

	fmt.Printf("%s %s\n", tableParam, values)

	err := Create(tableParam, values)

	return err
}

func ReadTransactions() ([]map[string]string, error) {
	transactionDetails := []map[string]string{}

	rows, err := ReadMultipleRows(tableTransactionDetails.name, tableTransactionDetails.fields)
	if err != nil {
		return nil, fmt.Errorf("Error querying table: %w", err)
	}

	for rows.Next() {
		var name, budget, account, amount string
		var date time.Time

		rows.Scan(&name, &amount, &budget, &account, &date)

		year, month, day := date.Date()
		modifiedDate := fmt.Sprintf("%d-%v-%v", month, day, year)

		valueMap := map[string]string{
			"Name": name,
			"Amount": amount,
			"Budget": budget,
			"Account": account,
			"Date": modifiedDate,
		}

		transactionDetails = append(transactionDetails, valueMap)
	}

	return transactionDetails, nil
}
