package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/auth"
)

func GetLoginUserAccountId(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.AdminJwtClaims)
	return claims.ID
}
