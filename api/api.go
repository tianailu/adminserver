package api

import (
	"github.com/tianailu/adminserver/api/admin/auth"
	"github.com/tianailu/adminserver/api/admin/user"
)

func InitTable() {
	auth.InitTable()
	user.InitTable()
}
