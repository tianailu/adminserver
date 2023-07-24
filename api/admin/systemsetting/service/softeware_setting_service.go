package service

import (
	"log"

	"github.com/tianailu/adminserver/api/admin/systemsetting/domain"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm"
)

func NewSoftwareSettingService() SoftwareSettingService {
	return SoftwareSettingService{
		db: mysql.GetDb(),
	}
}

type SoftwareSettingService struct {
	db *gorm.DB
}

func (svc *SoftwareSettingService) AddOrUpdateSoftwareSetting(softwareSetting *domain.SoftwareSetting) error {
	err := svc.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(softwareSetting).Error
	})
	if err != nil {
		log.Println("保存失败")
		return err
	}
	return nil
}

func (svc *SoftwareSettingService) GetSoftwareSetting(settingType string) (c string, e error) {
	var content string
	err := svc.db.
		Model(&domain.SoftwareSetting{}).
		Select("content").
		Where("type", settingType).
		Find(&content).Error
	return content, err
}
