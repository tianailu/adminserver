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

	g.GET("/users/:user_id", userController.FindUserDetail)
	g.GET("/users/audit", userController.FindUserSimpleListByAuditType)
	g.GET("/users", userController.FindUserList)
	g.POST("/users", userController.AddUser)
	g.PUT("/users/audit", userController.UpdateUserAuditStatus)

	g.GET("/users/uid", userController.CreateUid)

	friendController := controllers.NewFriendController()
	g.GET("/friends", friendController.FindFriendshipList)

	heartbeatController := controllers.NewHeartbeatController()
	g.GET("/heartbeats", heartbeatController.FindHeartbeatList)

	findCompanionController := controllers.NewFindCompanionController()
	g.GET("/companions", findCompanionController.FindCompanionList)
	g.GET("/companions/types", findCompanionController.FindCompanionTypeList)
	g.POST("/companions/types", findCompanionController.AddCompanionType)
}
