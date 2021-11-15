package main

import (
	"fmt"
	"oop-bank/accounts" // go env -w GO111MODULE=off
	"oop-bank/client"
)

func main() {
	account1 := accounts.BankAccount{Number: 1, Owner: client.Client{Name: "John"}, Balance: 100}
	account2 := accounts.BankAccount{Number: 1, Owner: client.Client{Name: "Mary"}, Balance: 300}
	account1.Transfer(50, &account2)
	fmt.Println(account1)
	fmt.Println(account2)
}
