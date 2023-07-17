package domain

type Role struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
	CreateUserId int64  `json:"createUserId"`
	UpdateUserId int64  `json:"updateUserId"`
}

func (Role) TableName() string {
	return "tb_role"
}

/*
SQL

-- tal.tb_role definition

CREATE TABLE `tb_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `create_time` bigint(20) DEFAULT NULL,
  `update_time` bigint(20) DEFAULT NULL,
  `create_user_id` bigint(20) NOT NULL,
  `update_user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `tb_role_un_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

*/
