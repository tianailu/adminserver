package service

import (
	"fmt"
	"github.com/tianailu/adminserver/api/admin/system_setting/domain"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm"
	"time"
)

func NewSoftwareSettingService() SoftwareSettingService {
	return SoftwareSettingService{
		db: mysql.GetDB(),
	}
}

type SoftwareSettingService struct {
	db *gorm.DB
}

func (svc *SoftwareSettingService) AddOrUpdateSoftwareSetting(userId string, softwareSetting *domain.SoftwareSetting) error {
	err := svc.db.Transaction(func(tx *gorm.DB) error {
		var id int
		_ = tx.Model(&domain.SoftwareSetting{}).Select("id").Where("type", softwareSetting.Type).First(&id).Error
		fmt.Printf("old setting id %#v\n", id)
		if id == 0 {
			softwareSetting.CreateAccountId = userId
			softwareSetting.CreateAt = time.Now().Unix()
		} else {
			softwareSetting.Id = id
			softwareSetting.UpdateAccountId = userId
			softwareSetting.UpdateAt = time.Now().Unix()
		}
		return tx.Save(softwareSetting).Error

	})
	if err != nil {
		fmt.Printf("保存失败\n")
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
