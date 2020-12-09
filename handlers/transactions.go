package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ezeoleaf/bankapp/db"
	"github.com/ezeoleaf/bankapp/models"
	echo "github.com/labstack/echo/v4"
)

// GetTransactionsByUser godoc
// @Summary Returns list of Transactions
// @Description Returns list of transactions for an user ID
// @Tags transactions
// @Accept json
// @Produce json
// @Param userID path int true "user ID"
// @Success 200 {object} db.TransactionList
// @Failure 400 {object} int
// @Router /transactions/{userID} [get]
func GetTransactionsByUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _ := strconv.ParseInt(c.Param("userID"), 10, 64)
		response, status := models.GetTransactions(userID)

		return c.JSON(status, response)
	}
}

// PostTransaction godoc
// @Summary Register a transaction
// @Description Creates and attach a transaction to a user account and calculate the balance
// @Tags transactions
// @Accept json
// @Produce json
// @Param user_id body int true "user id"
// @Param type body string true "debit"
// @Param description body string false "description"
// @Param amount body float64 false "amount"
// @Param entity body string false "bank"
// @Success 200 {object} db.Transaction
// @Failure 400 {object} int
// @Router /transactions/{transID} [post]
func PostTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		transReq := db.TransactionRequest{}
		err := c.Bind(&transReq)
		if err != nil || !validateTransactionRequest(transReq) {
			return c.JSON(http.StatusBadRequest, db.TransactionResponse{Success: false, Error: "Invalid request", ErrorCode: http.StatusBadRequest})
		}

		transReq.Type = strings.ToLower(transReq.Type)

		response := models.PostTransaction(transReq)
		return c.JSON(response.ErrorCode, response)
	}
}

func validateTransactionRequest(r db.TransactionRequest) bool {
	if r.Amount == 0 {
		return false
	}

	if strings.ToLower(r.Type) != db.DebitTransaction && strings.ToLower(r.Type) != db.CreditTransaction {
		return false
	}

	if r.UserID < 1 {
		return false
	}

	return true
}

// GetTransaction godoc
// @Summary Returns a transaction
// @Description Returns a transaction information from an ID
// @Tags transactions
// @Accept json
// @Produce json
// @Param transID path int true "transaction id"
// @Success 200 {object} db.Transaction
// @Failure 400 {object} int
// @Router /transactions/{transID} [get]
func GetTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		transID, _ := strconv.ParseInt(c.Param("transID"), 10, 64)
		response, status := models.GetTransaction(transID)

		return c.JSON(status, response)
	}
}
