package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/config"
)

func init() {
	go createDatabase()
}

func InitRouter(settings *config.SettingsConfig, g *echo.Group) {
	g.POST("/login", AdminLoginHandler(settings))
}
