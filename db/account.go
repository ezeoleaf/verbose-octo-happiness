package db

import "errors"

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
	a := Account{ID: 1, Balance: 0, UserID: 0}
	Accounts[1] = a
}

func GetAccount(userID int64) (Account, error) {
	if _, ok := Accounts[userID]; ok {
		return Accounts[userID], nil
	}

	return Account{}, errors.New("No account for this user")
}
