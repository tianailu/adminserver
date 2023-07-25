package domain

import "time"

type SoftwareSettingRequest struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type SoftwareSetting struct {
	Id              int       `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null;comment:主键"`
	Content         string    `json:"content" gorm:"column:content;comment:内容"`
	Type            string    `json:"type" gorm:"column:type;size:32;comment:类型:about-us、user-agreement，user-privacy-policy"`
	CreateAccountId string    `json:"createAccountId" gorm:"column:create_account_id;size:32;comment:创建用户id"`
	UpdateAccountId string    `json:"updateAccountId" gorm:"column:update_account_id;size:32;comment:更新用户id"`
	CreateAt        time.Time `json:"createAt" gorm:"column:create_at;type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	UpdateAt        time.Time `json:"updateAt" gorm:"column:update_at;type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:更新时间"`
}

func (ss *SoftwareSetting) TableName() string {
	return "tb_software_setting"
}

