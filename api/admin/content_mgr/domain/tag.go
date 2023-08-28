package domain

type Tag struct {
	Id              int    `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null;comment:主键"`
	Name            string `json:"name" gorm:"column:name;unique;not null;comment:标签名称"`
	Status          int8   `json:"status" gorm:"column:status;not null;comment:标签状态"`
	CreateAccountId string `json:"createAccountId" gorm:"column:create_account_id;comment:创建用户id"`
	UpdateAccountId string `json:"updateAccountId" gorm:"column:update_account_id;comment:内容更新用户id"`
	CreateAt        int64  `json:"createAt" gorm:"column:create_at;not null;comment:创建时间"`
	UpdateAt        int64  `json:"updateAt" gorm:"column:update_at;not null;comment:更新时间"`
}

func (r *Tag) TableName() string {
	return "tb_tag"
}
