package domain

type RolePermission struct {
	Id              int    `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null;comment:主键"`
	RoleId          int    `json:"roleId" gorm:"column:role_id;not null;uniqueIndex:udx_role_id_perm_id;comment:角色id"`
	PermissionId    int    `json:"permissionId" gorm:"column:permission_id;not null;uniqueIndex:udx_role_id_perm_id;comment:权限id"`
	CreateAccountId string `json:"createAccountId" gorm:"column:create_account_id;comment:创建用户id"`
	CreateAt        int64  `json:"createAt" gorm:"column:create_at;not null;comment:创建时间"`
}

func (rp *RolePermission) TableName() string {
	return "tb_role_permission"
}
