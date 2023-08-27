package req

type SaveRolePermissionRequest struct {
	RoleId        int   `json:"roleId"`
	PermissionIds []int `json:"permissionIds"`
}

type RolePageRequest struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
}
