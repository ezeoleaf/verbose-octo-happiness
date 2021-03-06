package db

import (
	"errors"
	"time"
)

func (a *Account) updateBalance(t Transaction) {
	if t.Type == DebitTransaction {
		a.Balance -= t.Amount
	} else if t.Type == CreditTransaction {
		a.Balance += t.Amount
	}
}

func (a *Account) updateTransactions(t Transaction) {
	a.History = append(a.History, t)
}

func CreateAccount() {
	Accounts = make(map[int64]Account)
	a := Account{ID: 1, Balance: 5000, UserID: 0}
	t := Transaction{ID: 1, Type: "credit", Description: "Initial bank account", Amount: 5000, Date: time.Now(), Entity: "Olaf"}
	a.History = append(a.History, t)
	Accounts[1] = a
}

func GetAccount(userID int64) (Account, error) {
	if _, ok := Accounts[userID]; ok {
		return Accounts[userID], nil
	}

	return Account{}, errors.New("No account for this user")
}
