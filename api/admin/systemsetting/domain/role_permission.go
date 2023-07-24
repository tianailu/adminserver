package domain

type RolePermission struct {
	Id           int   `json:"id"`
	RoleId       int   `json:"roleId"`
	PermissionId int   `json:"permissionId"`
	UserId       int   `json:"userId"`
	CreateTime   int64 `json:"createTime"`
	UpdateTime   int64 `json:"updateTime"`
	CreateUserId int   `json:"createUserId"`
	UpdateUserId int   `json:"updateUserId"`
}

/*
-- tal.tb_role_permission definition

CREATE TABLE `tb_role_permission` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_id` bigint(20) NOT NULL,
  `permission_id` bigint(20) NOT NULL,
  `create_time` bigint(20) DEFAULT NULL,
  `create_user_id` bigint(20) DEFAULT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  `update_user_id` bigint(20) DEFAULT NULL,
  `user_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `tb_role_permission_un` (`role_Id`,`permission_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;
*/

func (rp *RolePermission) TableName() string {
	return "tb_role_permission"
}
