package accounts

import (
	"errors"
	"fmt"
	"oop-bank/client"
)

type SavingAccount struct {
	Owner             client.Client
	Number, Operation int
	balance           float64
}

func (account *SavingAccount) Deposit(value float64) error {
	if value > 0 {
		account.balance += value
		return nil
	}
	return errors.New("you can't, deposit a negative value")
}

func (account *SavingAccount) Withdraw(value float64) error {
	if value > 0 {
		account.balance -= value
		return nil
	}
	return errors.New("you can't withdraw a negative value")
}

func (account *SavingAccount) Transfer(value float64, destination *SavingAccount) error {
	if value < 0 || account.balance < value {
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

func (account SavingAccount) GetBalance() float64 {
	return account.balance
}

func (account SavingAccount) String() string {
	return fmt.Sprintf(
		`{"Account": %d, "Owner": "%s", "balance": %.2f}`,
		account.Number,
		account.Owner,
		account.balance,
	)
}
