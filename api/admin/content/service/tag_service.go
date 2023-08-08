package service

import (
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm"
)

type TagService struct {
	db *gorm.DB
}

func NewTagService() TagService {
	return TagService{
		db: mysql.GetDb(),
	}
}
