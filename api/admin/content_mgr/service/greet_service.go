package service

import (
	"github.com/tianailu/adminserver/api/admin/content_mgr/domain"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm"
)

type GreetService struct {
	db *gorm.DB
}

func NewGreetService() *GreetService {
	return &GreetService{
		db: mysql.GetDB(),
	}
}

func (svc *GreetService) SaveOrUpdateGreet(greet domain.Greet) error {
	return svc.db.Model(&domain.Greet{}).Save(greet).Error
}
