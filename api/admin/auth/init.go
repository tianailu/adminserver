package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func init() {
}

func InitTable() {
	go func() {
		err := createTable()
		if err != nil {
			log.Errorf("Failed to init table in auth module, error: %s", err)
		}
	}()
}

func InitRouter(g *echo.Group) {
	g.POST("/login", AdminLogin)
}

func InitAdminRouter(g *echo.Group) {
	g.GET("/accounts", GetAdminList)
	g.POST("/accounts", AddAdmin)
	g.PUT("/accounts", UpdateAdmin)
}
