package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tianailu/adminserver/pkg/server"
)

func main() {
	as := server.NewAdminServer()
	as.Initialize()
	as.Start()
}
