package handlers

import (
	"github.com/ezeoleaf/bankapp/models"
	echo "github.com/labstack/echo/v4"
)

// GetTransactions godoc
// @Summary Returns list of Transactions
// @Description Returns list of transactions for a user ID
// @Tags transactions
// @Accept json
// @Produce json
// @Param from query string false "Start date to search market pairs "
// @Param to query string false "End date to search market pairs "
// @Success 200 {object} models.PricesResponse
// @Failure 400 {object} int
// @Router /transactions [get]
func GetTransactions() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("userID")
		response, status := models.GetTransactions(userID)

		return c.JSON(status, response)
	}
}

func PostDebit() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}

func PostCredit() echo.HandlerFunc {
	return func(c echo.Context) error {

	}
}

func GetTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		transID := c.Param("transID")
		response, status := models.GetTransaction(transID)

		return c.JSON(status, response)
	}
}
