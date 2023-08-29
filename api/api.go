package api

import (
	"github.com/tianailu/adminserver/api/admin/auth"
	"github.com/tianailu/adminserver/api/admin/content_mgr"
	"github.com/tianailu/adminserver/api/admin/system_setting"
	"github.com/tianailu/adminserver/api/admin/user"
)

func InitTable() {
	auth.InitTable()
	user.InitTable()
	system_setting.InitTable()
	content_mgr.InitTable()
}
