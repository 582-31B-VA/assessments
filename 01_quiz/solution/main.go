package main

import "fmt"

// askName asks for a name to create an account and returns the name.
func askName() string {
	fmt.Println("Enter your name to create an account.")
	var n string
	fmt.Scanln(&n)
	return n
}

// askAction asks which action to take (deposit or withdraw) and returns
// the actions as well as the amount to deposit or withdraw.
func askAction() (string, int) {
	fmt.Println("Enter 'deposit' or 'withdraw' followed by an amount.")
	var action string
	var amount int
	fmt.Scanln(&action, &amount)
	return action, amount
}

func main() {
	name := askName()
	bank := newBank()
	account := bank.newAccount(name, 0)
	action, amount := askAction()
	var balance cents
	switch action {
	case "deposit":
		balance = account.deposit(cents(amount))
	case "withdraw":
		var err error
		balance, err = account.withdraw(cents(amount))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	fmt.Println("Balance:", balance)
}
