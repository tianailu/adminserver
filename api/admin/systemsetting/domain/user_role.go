package domain

type UserRole struct {
	Id           int   `json:"id"`
	RoleId       int   `json:"roleId"`
	CreateTime   int64 `json:"createTime"`
	CreateUserId int   `json:"createUserId"`
}

func (ur *UserRole) TableName() string {
	return "tb_user_role"
}
