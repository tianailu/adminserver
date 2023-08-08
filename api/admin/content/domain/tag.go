package domain

import "time"

type Tag struct {
	Id              int       `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null;comment:主键"`
	Name            string    `json:"name" gorm:"column:name;unique;not null;comment:标签名称"`
	Status          string    `json:"status" gorm:"column:status;unique;not null;comment:标签状态"`
	CreateAccountId string    `json:"createAccountId" gorm:"column:create_account_id;comment:创建用户id"`
	UpdateAccountId string    `json:"updateAccountId" gorm:"column:update_account_id;comment:内容更新用户id"`
	CreateAt        time.Time `json:"createAt" gorm:"column:create_at;type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	UpdateAt        time.Time `json:"updateAt" gorm:"column:update_at;type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:更新时间"`
}

func (r *Tag) TableName() string {
	return "tb_tag"
}
