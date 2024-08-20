package main

import (
	"fmt"
)

// Cents are cents.
type cents int

// A bank is a bank.
type bank struct {
	accounts []account
	nextId   int
}

// An account is a bank account.
type account struct {
	id      int
	owner   string
	balance cents
}

// newBank creates and returns a bank.
func newBank() bank {
	return bank{[]account{}, 1}
}

// newAccount creates an account for the given owner with the given
// balance, and stores it in the received bank.
func (b bank) newAccount(owner string, balance cents) account {
	a := account{b.nextId, owner, balance}
	b.nextId += 1
	b.accounts = append(b.accounts, a)
	return a
}

// withdraw withdraws money from the received account and returns the
// account's balance. If amount is greater than the account's balance,
// it returns an error.
func (a *account) withdraw(amount cents) (cents, error) {
	if a.balance < amount {
		return a.balance, fmt.Errorf(
			"withdrawing %v: insufficient funds, current balance: %v",
			amount,
			a.balance,
		)
	}
	a.balance -= amount
	return a.balance, nil
}

// deposit deposits money to the received account and returns the
// account's balance.
func (a *account) deposit(amount cents) cents {
	a.balance += amount
	return a.balance
}
