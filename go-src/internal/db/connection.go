package db

import (
	"os"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	var err error
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_STRING"))
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %w", err)
	}

	if conn.Ping(context.Background()) != nil {
		return nil, fmt.Errorf("Unable to ping database: %w", err)
	}

	fmt.Println("Connected!")

	return conn, nil
}

func Create(table string, values string) error {
	conn, err := Connect()
	if err != nil {
		return fmt.Errorf("Database error: %w", err)
	}

	defer conn.Close(context.Background())

	sql := fmt.Sprintf("INSERT INTO %s VALUES (%s)", table, values)

	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("Error creating values: %w", err)
	}

	log.Printf("Added %s values into table %s\n", values, table)
	return nil
}

func Delete(table string, condition string) error {
	conn, err := Connect()
	if err != nil {
		return fmt.Errorf("Database error: %w", err)
	}

	defer conn.Close(context.Background())

	sql := fmt.Sprintf("DELETE FROM %s WHERE (%s)", table, condition)

	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("Error deleting values: %w", err)
	}

	log.Printf("Deleted %s values from table %s\n", condition, table)
	return nil
}

func Update(table string, updatedValue string, nameToUpdate string) error {
	conn, err := Connect()
	if err != nil {
		return fmt.Errorf("Database error: %w", err)
	}

	defer conn.Close(context.Background())

	sql := fmt.Sprintf("UPDATE %s SET %s WHERE name = '%s'", table, updatedValue, nameToUpdate)

	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("Error updating values: %w", err)
	}

	log.Printf("Updated %s values in table %s\n", updatedValue, table)
	return nil
}

func Read(table string, field string, filter string) (pgx.Row, error) {
	conn, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("Database error: %w", err)
	}

	defer conn.Close(context.Background())

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s", field, table, filter)

	return conn.QueryRow(context.Background(), sql), nil
}
