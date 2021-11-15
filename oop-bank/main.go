package main

import (
	"fmt"
	"reflect"
)

func main() {
	// constructStructure()
	// withdrawExample()
	// depositExample()
	// transferExample()
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

func transferExample() {
	account1 := BankAccount{owner: "John", agency: 1, account: 1, balance: 500}
	account2 := BankAccount{owner: "Jane", agency: 1, account: 2, balance: 500}

	account1.transfer(100, &account2)
	fmt.Println(account1.balance) // 400
	fmt.Println(account2.balance) // 600
	fmt.Println(account1, account2)
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

func (account *BankAccount) withdraw(value float64) *AccountError {
	if value > 0 {
		account.balance -= value
		return nil
	}
	return &AccountError{"Error, you can't withdraw a negative value"}
}

func (account *BankAccount) deposit(value float64) *AccountError {
	if value > 0 {
		account.balance += value
		return nil
	}
	return &AccountError{"Error, you can't deposit a negative value"}
}

type AccountError struct {
	ErrMsg string
}

func (e *AccountError) Error() string {
	return e.ErrMsg
}

// variadic arguments
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func (account *BankAccount) transfer(value float64, account2 *BankAccount) *AccountError {
	if value > account.balance && value > 0 {
		return &AccountError{"Error, account balance is not enough"}
	}
	err := account.deposit(value)
	if err != nil {
		return err
	}
	err = account2.withdraw(value)
	if err != nil {
		return err
	}
	return nil
}

// *pointer
// &address
