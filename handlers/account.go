package handlers

import (
	"strconv"

	"github.com/ezeoleaf/bankapp/models"
	echo "github.com/labstack/echo/v4"
)

// GetAccount godoc
// @Summary Returns account for a user
// @Description Returns an account from an user ID
// @Tags transactions
// @Accept json
// @Produce json
// @Param userID path int true "user ID"
// @Success 200 {object} db.Account
// @Failure 400 {object} int
// @Router /account/{userID} [get]
func GetAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _ := strconv.ParseInt(c.Param("userID"), 10, 64)
		response, status := models.GetAccount(userID)

		return c.JSON(status, response)
	}
}
