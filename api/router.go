package api

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/auth"
	"github.com/tianailu/adminserver/api/admin/user"
	"github.com/tianailu/adminserver/config"
)

// InitRouter 初始化基本路由
func InitRouter(e *echo.Echo) {
}

func InitAdminRouter(setting *config.SettingsConfig, g *echo.Group) {
	auth.InitRouter(setting, g)
}

func InitGroupAdminRouter(ver *echo.Group) {
	ver.POST("/user", user.AddUser)
	ver.GET("/user", user.GetUser)
	ver.GET("/user/vip", user.GetVipUser)
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
