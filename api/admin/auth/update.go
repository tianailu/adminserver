package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpdateAdmin(c echo.Context) error {

	return c.JSON(http.StatusOK, nil)
}
