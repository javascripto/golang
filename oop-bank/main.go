package main

import (
	"fmt"
	"oop-bank/accounts" // go env -w GO111MODULE=off
	"oop-bank/client"
)

func main() {
	account1 := accounts.BankAccount{Number: 1, Owner: client.Client{Name: "John"}}
	account2 := accounts.BankAccount{Number: 1, Owner: client.Client{Name: "Mary"}}
	account3 := accounts.SavingAccount{Number: 1, Owner: client.Client{Name: "Mary"}}
	account1.Deposit(500)
	account2.Deposit(500)
	account3.Deposit(200)
	account1.Transfer(50, &account2)
	fmt.Println(account1)
	fmt.Println(account2)

	PayBankBillet(&account3, 100)
	fmt.Println(account3)
}

func PayBankBillet(account VerifyAccount, bankBilletValue float64) error {
	return account.Withdraw(bankBilletValue)
}

type VerifyAccount interface {
	Withdraw(value float64) error
}
