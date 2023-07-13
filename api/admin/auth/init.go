package auth

import (
	"github.com/labstack/echo/v4"
)

func init() {
	go createDatabase()
}

func InitRouter(g *echo.Group) {
	g.POST("/login", AdminLogin)
}
