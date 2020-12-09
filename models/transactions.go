package models

import (
	"net/http"

	"github.com/ezeoleaf/bankapp/db"
)

func GetTransactions(userID int64) (db.TransactionList, int) {
	return db.GetTransactionsByUser(userID), http.StatusOK
}

func GetTransaction(transID int64) (db.Transaction, int) {
	return db.GetTransactionByID(transID), http.StatusOK
}

func PostTransaction(r db.TransactionRequest) db.TransactionResponse {
	return db.SaveTransactionRequest(r)
}
