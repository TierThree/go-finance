-- Create the Budgets table
-- ID: integer, auto-incrementing serial primary key
-- Name: text, indicates the name of the budget
-- Amount: decimal, indicates the amount of the budget

CREATE TABLE budgets (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	amount decimal NOT NULL
);

-- Create the Accounts table
-- ID: integer, auto-incrementing serial primary key
-- Name: text, indicates the name of the account
-- Balance: decimal, indicates the balance of the account

CREATE TABLE accounts (
	id SERIAL PRIMARY KEY,
	name text NOT NULL,
	balance decimal NOT NULL
);

-- Create the transactions table
-- ID: integer, auto-incrementing serial primary key
-- name: text, name of the transaction
-- amount: decimal, dollar amount of the transaction
-- budget: text, name of the budget used
-- account: text, name of the account used
-- date: date, date that the transaction occurred

CREATE TABLE transactions (
	id SERIAL PRIMARY KEY,
	name text NOT NULL,
	amount decimal NOT NULL,
	budget text NOT NULL,
	account text NOT NULL,
	date date NOT NULL
);
