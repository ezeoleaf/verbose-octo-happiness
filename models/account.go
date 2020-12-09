package models

import (
	"net/http"

	"github.com/ezeoleaf/bankapp/db"
)

func GetAccount(userID int64) (db.Account, int) {
	a, e := db.GetAccount(userID)
	if e != nil {
		return a, http.StatusBadRequest
	}

	return a, http.StatusOK
}
