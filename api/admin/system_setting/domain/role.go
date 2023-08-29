package domain

type Role struct {
	Id              int    `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null;comment:主键"`
	Name            string `json:"name" gorm:"column:name;unique;not null;comment:角色名称"`
	CreateAccountId string `json:"createAccountId" gorm:"column:create_account_id;comment:创建用户id"`
	CreateAt        int64  `json:"createAt" gorm:"column:create_at;not null;comment:创建时间"`
}

func (r *Role) TableName() string {
	return "tb_role"
}
