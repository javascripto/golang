package main

import (
	"fmt"
	"oop-bank/accounts" // go env -w GO111MODULE=off
)

func main() {
	account1 := accounts.BankAccount{Number: 1, Owner: "John", Balance: 100}
	account2 := accounts.BankAccount{Number: 1, Owner: "Mary", Balance: 300}
	account1.Transfer(50, &account2)
	fmt.Println(account1)
	fmt.Println(account2)
}
