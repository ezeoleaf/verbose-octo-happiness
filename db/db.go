package db

import "time"

type Account struct {
	ID      int64         `json:"id"`
	Balance float64       `json:"balance"`
	UserID  int64         `json:"user_id"`
	History []Transaction `json:"history"`
}

type Transaction struct {
	ID          int64     `json:"id"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Entity      string    `json:"entity"`
}

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TransactionRequest struct {
	UserID      int64   `json:"user_id"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Entity      string  `json:"entity"`
}

type TransactionResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

var Accounts map[int64]Account
var Transactions map[int64]Transaction

type TransactionList struct {
	Transactions []Transaction `json:"transactions"`
}

const (
	DebitTransaction  = "debit"
	CreditTransaction = "credit"
)
