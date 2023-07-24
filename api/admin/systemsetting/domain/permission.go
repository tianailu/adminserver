package domain

type Permission struct {
	Id           int    `json:"id" gorm:"column:id"`
	Name         string `json:"name" gorm:"column:name"`
	ParentId     int    `json:"parentId" gorm:"column:parent_id"`
	Route        string `json:"route" gorm:"column:route"`
	Order        int    `json:"order"`
	CreateTime   int64  `json:"createTime" gorm:"column:create_time"`
	CreateUserId int    `json:"createUserId" gorm:"column:create_user_id"`
}

func (p *Permission) TableName() string {
	return "tb_permission"
}

/*
SQL
CREATE TABLE tal.tb_permission (
	id BIGINT auto_increment NOT NULL,
	name VARCHAR(50) NOT NULL,
	create_time BIGINT NOT NULL,
	route VARCHAR(200) NOT NULL,
	create_user_id BIGINT NULL,
	parent_id BIGINT NULL,
	order int NULL,
	CONSTRAINT tb_permission_pk PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;



data


INSERT INTO tal.tb_permission (name,create_time,route,create_user_id,parent_id) VALUES
	 ('用户管理',NULL,'/user-management',NULL,0),
	 ('内容管理',NULL,'/c',NULL,0),
	 ('聊天管理',NULL,'/chat-management',NULL,0),
	 ('用户列表',NULL,'/user-managerment/user-list',NULL,1),
	 ('举报管理',NULL,'/user-management/report',NULL,1),
	 ('禁言管理',NULL,'/user/management/ban',NULL,1),
	 ('封号管理',NULL,'/user/management/ban-user',NULL,1),
	 ('发帖列表',NULL,'/c/post-list',NULL,2),
	 ('评论列表',NULL,'/c/coment-list',NULL,2),
	 ('聊天设置',NULL,'/chat-managerment/chat-setting',NULL,3);

*/
