package accounts

import (
	"errors"
	"fmt"
	"oop-bank/client"
)

type BankAccount struct {
	Number  int
	Owner   client.Client
	Balance float64
}

func (account *BankAccount) Deposit(value float64) error {
	if value > 0 {
		account.Balance += value
		return nil
	}
	return errors.New("you can't, deposit a negative value")
}

func (account *BankAccount) Withdraw(value float64) error {
	if value > 0 {
		account.Balance -= value
		return nil
	}
	return errors.New("you can't withdraw a negative value")
}

func (account *BankAccount) Transfer(value float64, destination *BankAccount) error {
	if value < 0 || account.Balance < value {
		return errors.New("account balance is not enough")
	}
	err := account.Withdraw(value)
	if err != nil {
		return err
	}
	err = destination.Deposit(value)
	if err != nil {
		return err
	}
	return nil
}

// String method needs to reference the struct value, not its pointer
func (account BankAccount) String() string {
	return fmt.Sprintf(
		`{"Account": %d, "Owner": "%s", "Balance": %.2f}`,
		account.Number,
		account.Owner,
		account.Balance,
	)
}
