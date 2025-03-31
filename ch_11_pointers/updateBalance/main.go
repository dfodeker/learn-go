package main

import "fmt"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// implement the updateBalance function. It should take a customer pointer and a transaction,
// and return an error. Depending on the transactionType, it should either add or subtract the
//  amount from the customers balance. If the customer does not have enough money,
// it should return the error insufficient funds. If the transactionType isn't a withdrawal or deposit,
// it should return the error unknown transaction type. Otherwise, it should process the transaction and return nil.

// alice := customer{id: 1, balance: 100.0}
// deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

// updateBalance(&alice, deposit)
// id: 1 balance: 150
// Don't touch above this line

// ?
func updateBalance(customer *customer, transaction transaction) error {
	//depending on the transaction type , it should either add or subtract the amount from the account balance///

	//if the customer does not have enough money it should return error insufficient funds error

	if customer.balance < transaction.amount {
		return fmt.Errorf("insufficient funds")
	}
	if transaction.transactionType != "deposit" && transaction.transactionType != "withdrawal" {
		return fmt.Errorf("unknown transaction type")
	}

	if transaction.transactionType == "deposit" {
		customer.balance += transaction.amount
	}
	if transaction.transactionType == "withdrawal" {
		customer.balance -= transaction.amount
	}

	return nil
}
