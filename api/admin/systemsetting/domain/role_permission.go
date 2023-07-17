package domain

type RolePermission struct {
	Id           int   `json:"id"`
	RoleId       int   `json:"roleId"`
	PermissionId int   `json:"permissionId"`
	CreateTime   int64 `json:"createTime"`
	UpdateTime   int64 `json:"updateTime"`
	CreateUserId int   `json:"createUserId"`
	UpdateUserId int   `json:"updateUserId"`
}

/*
CREATE TABLE tal.tb_role_permission (
	id BIGINT auto_increment NOT NULL,
	roleId BIGINT NOT NULL,
	permission_id BIGINT NOT NULL,
	create_time BIGINT NULL,
	create_user_id BIGINT NULL,
	update_time BIGINT NULL,
	update_user_id BIGINT NULL,
	CONSTRAINT tb_role_permission_pk PRIMARY KEY (id),
	CONSTRAINT tb_role_permission_un UNIQUE KEY (roleId,permission_id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;

*/

func (RolePermission) TableName() string {
	return "tb_role_permission"
}
