package user

import "github.com/labstack/gommon/log"

func init() {
	go dealUserData()
}

func InitTable() {
	go func() {
		err := createTable()
		if err != nil {
			log.Errorf("Failed to init table in user module, error: %s", err)
		}
	}()
}
