package db

import "time"

type Account struct {
	ID      int64         `json:"id"`
	Balance float64       `json:"balance"`
	UserID  int64         `json:"user_id"`
	History []Transaction `json:"history"`
}

type Transaction struct {
	ID          int64           `json:"id"`
	Type        TransactionType `json:"type"`
	Description string          `json:"description"`
	Amount      float64         `json:"amount"`
	Date        time.Time       `json:"date"`
}

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TransactionType string

var Accounts map[int64]Account
