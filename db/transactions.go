package db

import (
	"fmt"
	"time"
)

func GetTransactionsByUser(userID int64) TransactionList {
	if _, ok := Accounts[userID]; ok {
		fmt.Println(Accounts)
		return TransactionList{Transactions: Accounts[userID].History}
	}
	fmt.Println(Accounts)

	return TransactionList{}
}

func GetTransactionByID(transID int64) Transaction {
	if _, ok := Transactions[transID]; ok {
		return Transactions[transID]
	}

	return Transaction{}
}

func SaveTransactionRequest(r TransactionRequest) TransactionResponse {
	tr := TransactionResponse{Success: false}
	if _, ok := Accounts[r.UserID]; !ok {
		tr.Error = "There is no bank account for this user"
		tr.ErrorCode = 400
		return tr
	}

	account := Accounts[r.UserID]

	if r.Type == DebitTransaction {
		if r.Amount > account.Balance {
			tr.Error = "There is no enought balance in this account"
			tr.ErrorCode = 401
			return tr
		}
	}

	t := createTransaction(r)
	fmt.Println(t)
	account.updateBalance(t)
	account.updateTransactions(t)

	fmt.Println(account)
	Accounts[r.UserID] = account
	tr.Success = true

	return tr
}

func createTransaction(r TransactionRequest) Transaction {
	nextReq := int64(len(Transactions) + 1)

	t := Transaction{
		ID:          nextReq,
		Type:        r.Type,
		Description: r.Description,
		Amount:      r.Amount,
		Date:        time.Now(),
		Entity:      r.Entity,
	}

	Transactions[nextReq] = t

	return t
}

func CreateTransactions() {
	Transactions = make(map[int64]Transaction)
	t := Transaction{ID: 1, Type: "credit", Description: "Initial bank account", Amount: 5000, Date: time.Now(), Entity: "Olaf"}
	Transactions[1] = t
}
