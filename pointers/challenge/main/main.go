/*
                 *** Update Balance ***
Textio needs a new way to update users account balance.

Implement the updateBalance function. It should take
a pointer to a customer and a transaction,
and return an error. Depending on the transactionType,
it should either add or subtract the amount
from the customers balance. If the customer does not have
enough money, it should return the error insufficient funds.
If the transactionType isn't a withdrawal or deposit,
it should return the error unknown transaction type.
Otherwise, it should process the transaction and return nil.

alice := customer{id: 1, balance: 100.0}
deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

updateBalance(&alice, deposit)
// id: 1 balance: 150

*/

package main

import (
	"errors"
)

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

// Don't touch above this line

func updateBalance(cPtr *customer, t transaction) error {

    if t.transactionType == transactionDeposit {
        cPtr.balance += t.amount
        return nil
    }

    if t.transactionType == transactionWithdrawal {
        if cPtr.balance < t.amount {
            return errors.New("insufficient funds")
        }
        cPtr.balance -= t.amount
        return nil
    }

    return errors.New("unknown transaction type")
}
