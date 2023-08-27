package domain

import "time"

type UserRole struct {
	Id              int       `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null;comment:主键"`
	AccountId       int       `json:"accountId" gorm:"column:account_id;not null;comment:账户id"`
	RoleId          int       `json:"roleId" gorm:"column:role_id;not null;comment:角色id"`
	CreateAt        time.Time `json:"createAt" gorm:"column:create_at;type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	CreateAccountId int       `json:"createAccountId" gorm:"column:create_account_id;not null;comment:创建账户id"`
}

func (ur *UserRole) TableName() string {
	return "tb_user_role"
}
