package api

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/auth"
	"github.com/tianailu/adminserver/api/admin/gold"
	"github.com/labstack/echo"
	"github.com/tianailu/adminserver/api/admin/systemsetting"
	"github.com/tianailu/adminserver/api/admin/user"
)

// InitRouter 初始化基本路由
func InitRouter(e *echo.Echo) {
}

func InitAdminRouter(g *echo.Group) {
	auth.InitRouter(g)
	gold.InitRouter(g)
}

func InitGroupAdminRouter(g *echo.Group) {
	auth.InitAdminRouter(g)
	user.InitAdminRouter(g)
}

func InitSystemSettingRouter(e *echo.Echo) {
	// type支持 about-us、user-agreement，user-privacy-policy
	e.POST("/system-setting/:type", systemsetting.AddOrUpdateSoftwareSetting)
	e.GET("/system-setting/:type", systemsetting.GetSoftwareSetting)
	e.POST("/system-setting/roles", systemsetting.SaveRole)
	e.DELETE("/system-setting/roles/:id", systemsetting.DeleteRole)
	e.GET("/system-setting/roles", systemsetting.GetAllRoles)
	e.POST("/system-setting/roles/page", systemsetting.GetRolesPage)
	e.GET("/system-setting/permissions", systemsetting.GetUserPermisions)
	e.POST("/system-setting/role-permissions", systemsetting.SaveRolePermissions)
}

// 所有HTTP方法的接口，主要用于简化类型转换
type (
	GET interface {
		Get(c echo.Context) error
	}
	POST interface {
		Post(c echo.Context) error
	}
	PUT interface {
		Put(c echo.Context) error
	}
	DELETE interface {
		Delete(c echo.Context) error
	}

	HEAD interface {
		Head(c echo.Context) error
	}

	PATCH interface {
		Patch(c echo.Context) error
	}

	OPTIONS interface {
		Options(c echo.Context) error
	}
)

// Resource 资源
type Resource struct {
	Name string
}

func (r *Resource) GetName() string {
	return r.Name
}

func (r *Resource) SetName(name string) error {
	r.Name = name
	return nil
}
