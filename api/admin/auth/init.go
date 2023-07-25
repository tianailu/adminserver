package auth

import (
	"github.com/labstack/echo/v4"
)

func init() {
}

func InitTable() {
	go createTable()
}

func InitRouter(g *echo.Group) {
	g.POST("/login", AdminLogin)
}

func InitAdminRouter(g *echo.Group) {
	g.GET("/accounts", GetAdminList)
	g.POST("/accounts", AddAdmin)
	g.PUT("/accounts", UpdateAdmin)
}
