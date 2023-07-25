package systemsetting

import (
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/tianailu/adminserver/api/admin/systemsetting/controller"
	"github.com/tianailu/adminserver/api/admin/systemsetting/domain"
	"github.com/tianailu/adminserver/pkg/db/mysql"
)

func init() {
	go initSystemSettingData()
}

func initSystemSettingData() error {
	time.Sleep(time.Second * 5)
	log.Printf("begin init system setting table and data")

	tbSoftwareSetting := domain.SoftwareSetting{}
	err := mysql.GetDb().AutoMigrate(&tbSoftwareSetting)
	if err != nil {
		log.Printf("create system setting table[%s] failed,err: %s", tbSoftwareSetting.TableName(), err)
		return err
	}

	tbRole := domain.Role{}
	err = mysql.GetDb().AutoMigrate(&tbRole)
	if err != nil {
		log.Printf("create system setting table[%s] failed,err: %s", tbRole.TableName(), err)
		return err
	}

	tbUserRole := domain.UserRole{}
	err = mysql.GetDb().AutoMigrate(&tbUserRole)
	if err != nil {
		log.Printf("create system setting table[%s] failed,err: %s", tbUserRole.TableName(), err)
		return err
	}

	tbPermission := domain.Permission{}
	err = mysql.GetDb().AutoMigrate(&tbPermission)
	if err != nil {
		log.Printf("create system setting table[%s] failed,err: %s", tbPermission.TableName(), err)
		return err
	}

	tbRolePermission := domain.RolePermission{}
	err = mysql.GetDb().AutoMigrate(&tbRolePermission)
	if err != nil {
		log.Printf("create system setting table[%s] failed,err: %s", tbRolePermission.TableName(), err)
		return err
	}

	return nil
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
