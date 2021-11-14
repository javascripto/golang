package main

import (
	"fmt"
)

func main() {
	bankAccount1 := BankAccount{"John", 1, 1, 0}
	bankAccount2 := BankAccount{owner: "John", account: 1, agency: 1, balance: 0}
	bankAccounts := []BankAccount{bankAccount1, bankAccount2}
	fmt.Println(bankAccounts)
	fmt.Println(bankAccount2.toJSON())
}

type BankAccount struct {
	owner   string
	agency  int
	account int
	balance float64
}

func (account BankAccount) toJSON() string {
	return `{
  "owner": "` + account.owner + `",
  "accountNumber": ` + fmt.Sprintf("%d", account.account) + `,
  "agency": ` + fmt.Sprintf("%d", account.agency) + `,
  "balance": ` + fmt.Sprintf("%.2f", account.balance) + `
}`
}
