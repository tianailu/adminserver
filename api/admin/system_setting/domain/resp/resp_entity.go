package resp

type PermissionResponse struct {
	Id       int                   `json:"id"`
	Name     string                `json:"name"`
	Route    string                `json:"route"`
	ParentId int                   `json:"parentId"`
	Sequence int                   `json:"order"`
	Child    []*PermissionResponse `json:"child" gorm:"-"`
}

type RolePermissionDetail struct {
	Id       int                     `json:"id"`
	Name     string                  `json:"name"`
	Route    string                  `json:"route"`
	ParentId int                     `json:"parentId"`
	Sequence int                     `json:"order"`
	Enable   bool                    `json:"enable" gorm:"-"` // 当parentId == 0 时，默认为true
	Child    []*RolePermissionDetail `json:"child" gorm:"-"`
}

type UserRolePermissions struct {
	RoleId       int                     `json:"roleId"`
	RoleName     string                  `json:"roleName"`
	RoleCreateAt int64                   `json:"roleCreateAt"`
	Permissions  []*RolePermissionDetail `json:"permissions"`
}
