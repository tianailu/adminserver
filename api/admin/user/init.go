package user

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tianailu/adminserver/api/admin/user/controllers"
	"github.com/tianailu/adminserver/api/admin/user/models"
)

func init() {
	go dealUserData()
}

func InitTable() {
	go func() {
		err := models.CreateTable()
		if err != nil {
			log.Errorf("Failed to init table in user module, error: %s", err)
		}
	}()
}

func InitAdminRouter(g *echo.Group) {
	userController := controllers.NewUserController()

	g.GET("/users/:uid", userController.FindUserDetail)
	g.POST("/users", userController.AddUser)

	g.GET("/users/uid", userController.CreateUid)
}
