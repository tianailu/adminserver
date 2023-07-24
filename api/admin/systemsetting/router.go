package systemsetting

import (
	"github.com/labstack/echo"
	"github.com/tianailu/adminserver/api/admin/systemsetting/controller"
)

func init() {

}

func InitRouter(e *echo.Group) {
}

func InitGroupAdminRouter(e *echo.Group) {

	ctl := controller.NewSoftwareSettingController()
	// type支持 about-us、user-agreement，user-privacy-policy
	e.POST("/system-setting/:type", ctl.AddOrUpdateSoftwareSetting)
	e.GET("/system-setting/:type", ctl.GetSoftwareSetting)
	roleCtl := controller.NewRoleController()
	e.POST("/system-setting/roles", roleCtl.SaveRole)
	e.DELETE("/system-setting/roles/:id", roleCtl.DeleteRole)
	e.GET("/system-setting/roles", roleCtl.GetAllRoles)
	e.GET("/system-setting/role/:roleId/permissions", roleCtl.GetRolePermissions)
	e.POST("/system-setting/roles/page", roleCtl.GetRolesPage)
	e.POST("/system-setting/role-permissions", roleCtl.SaveRolePermissions)

	// perCtl := controller.NewPermissionController()
	// e.GET("/system-setting/permissions", perCtl.GetUserPermisions)
}
