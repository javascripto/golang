package main

import (
	"fmt"
	"reflect"
)

func main() {
	// constructStructure()
	// withdrawExample()
	// depositExample()
}

func constructStructure() {
	var bankAccount1 *BankAccount = new(BankAccount)                              // ponting to a struct
	var bankAccount2 *BankAccount = &BankAccount{}                                // dereferencing to a struct
	bankAccount3 := BankAccount{"John", 1, 1, 0}                                  // struct literal
	bankAccount4 := BankAccount{owner: "John", account: 1, agency: 1, balance: 0} // struct literal with field names

	bankAccounts := []BankAccount{ // slice of structs
		*bankAccount1,
		*bankAccount2,
		bankAccount3,
		bankAccount4,
	}

	fmt.Println(bankAccounts)
	fmt.Println(bankAccount1)
	fmt.Println(bankAccount3)
	fmt.Println(&bankAccount1) // memory address of reference
	fmt.Println(&bankAccount3) // reference to a struct
	fmt.Println(bankAccount2.toJSON())
	fmt.Println(bankAccount4.toJSON())
	fmt.Println(bankAccount1 == bankAccount2)   // false because they are different pointers
	fmt.Println(*bankAccount1 == *bankAccount2) // true because they are the same struct
	fmt.Println(&bankAccount3 == &bankAccount4) // false because they are different pointers

	reference := &bankAccount3
	memoryAddress := &reference

	fmt.Println(reference)                     // &{John 1 1 0}
	fmt.Println(memoryAddress)                 // 0xc00000e038
	fmt.Println(reflect.TypeOf(reference))     // *main.BankAccount
	fmt.Println(reflect.TypeOf(memoryAddress)) // **main.BankAccount
	fmt.Println(fmt.Sprintf("%p", reference))  // 0xc00000e038 as string
}

func withdrawExample() {
	account := BankAccount{"John", 1, 1, 500}
	fmt.Println(account)

	// function call with arguments by value
	account2 := WithdrawByValue(account, 100)
	fmt.Println(account.balance)  // 500
	fmt.Println(account2.balance) // 400

	// function call with arguments by pointer
	WithdrawByPointer(&account, 100)
	fmt.Println(account.balance) // 400 (500 - 100)

	// method call by pointer with arguments
	account.WithdrawByPointer(100)
	fmt.Println(account.balance) // 300 (400 - 100)

	// method call by value with arguments
	account3 := account.WithdrawByValue(100)
	fmt.Println(account.balance == account3.balance) // false
	fmt.Println(account.balance)                     // 300 (method call by value returns a copy of the struct)
	fmt.Println(account3.balance)                    // 200 (copy of the struct changes its value)
}

func depositExample() {
	account := BankAccount{"John", 1, 1, 500}

	account.deposit(100)
	err := account.deposit(-100)

	if err != nil {
		fmt.Println(err.ErrMsg)
	}

	fmt.Println(account.balance) // 600
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

func WithdrawByValue(account BankAccount, value float64) BankAccount {
	account.balance -= value
	return account
}

func WithdrawByPointer(account *BankAccount, value float64) {
	account.balance -= value
}

func (account *BankAccount) WithdrawByPointer(value float64) {
	account.balance -= value
}

func (account BankAccount) WithdrawByValue(value float64) BankAccount {
	account.balance -= value
	return account
}

func (account *BankAccount) deposit(value float64) *DepositError {
	if value > 0 {
		account.balance += value
		return nil
	}
	return &DepositError{"Error, you can't deposit a negative value"}
}

type DepositError struct {
	ErrMsg string
}

// variadic arguments
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
