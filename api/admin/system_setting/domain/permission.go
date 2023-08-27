package domain

import "time"

type Permission struct {
	Id              int       `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null;comment:主键"`
	Name            string    `json:"name" gorm:"column:name;unique;not null;uniqueIndex:uqIdx_name"`
	ParentId        int       `json:"parentId" gorm:"column:parent_id;not null"`
	Route           string    `json:"route" gorm:"column:route"`
	Sequence        int       `json:"sequence" gorm:"column:sequence"`
	CreateAt        time.Time `json:"createAt" gorm:"column:create_at;type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	CreateAccountId string    `json:"createAccountId" gorm:"column:create_account_id;size:32"`
}

func (p *Permission) TableName() string {
	return "tb_permission"
}

/*
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (1, '用户管理', 0, NULL, 1, '2023-07-26 20:45:33', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (2, '用户列表', 1, NULL, 2, '2023-07-26 20:45:33', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (3, '举报管理', 1, NULL, 3, '2023-07-26 20:46:53', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (4, '禁言管理', 1, NULL, 4, '2023-07-26 20:47:02', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (5, '封号管理', 1, NULL, 5, '2023-07-26 20:47:14', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (6, '冻结管理', 1, NULL, 6, '2023-07-26 20:47:22', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (7, '注销列表', 1, NULL, 7, '2023-07-26 20:47:34', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (8, '交友管理-用户列表', 1, NULL, 8, '2023-07-26 20:47:59', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (9, '心动列表', 1, NULL, 9, '2023-07-26 20:48:07', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (10, '标签管理', 1, NULL, 10, '2023-07-26 20:48:12', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (11, '打招呼管理', 1, NULL, 11, '2023-07-26 20:48:21', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (12, '用户详情', 1, NULL, 12, '2023-07-26 20:48:29', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (13, 'VIP等级', 1, NULL, 13, '2023-07-26 20:48:34', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (14, 'VIP列表', 1, NULL, 14, '2023-07-26 20:48:42', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (15, '交友设置', 1, NULL, 15, '2023-07-26 20:48:50', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (16, '交友主页设置', 1, NULL, 16, '2023-07-26 20:48:55', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (17, '推荐页设置', 1, NULL, 17, '2023-07-26 20:49:02', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (18, '资料编辑页设置', 1, NULL, 18, '2023-07-26 20:49:12', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (19, '用户体验设置', 1, NULL, 19, '2023-07-26 20:49:17', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (20, '消息设置', 1, NULL, 20, '2023-07-26 20:49:22', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (21, '收费设置', 1, NULL, 21, '2023-07-26 20:49:27', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (22, '内容管理', 0, NULL, 22, '2023-07-26 20:49:33', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (23, '发帖列表', 22, NULL, 23, '2023-07-26 20:49:42', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (24, '评论列表', 22, NULL, 24, '2023-07-26 20:49:52', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (25, '飞鸽传书', 22, NULL, 25, '2023-07-26 20:50:01', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (26, '共鸣', 22, NULL, 26, '2023-07-26 20:50:09', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (27, '聊天管理', 0, NULL, 27, '2023-07-26 20:50:16', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (28, '聊天设置', 27, NULL, 28, '2023-07-26 20:50:23', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (29, '金币管理', 0, NULL, 29, '2023-07-26 20:50:40', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (30, '金币记录', 29, NULL, 30, '2023-07-26 20:50:47', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (31, '金币银行', 29, NULL, 31, '2023-07-26 20:50:53', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (32, '资料记录', 29, NULL, 32, '2023-07-26 20:51:02', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (33, '签到记录', 29, NULL, 33, '2023-07-26 20:51:12', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (34, '邀请记录', 29, NULL, 34, '2023-07-26 20:51:19', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (35, '邀请记录码', 29, NULL, 35, '2023-07-26 20:51:30', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (36, '金币设置', 29, NULL, 36, '2023-07-26 20:51:37', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (37, '资料设置', 29, NULL, 37, '2023-07-26 20:51:47', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (38, '分享记录', 29, NULL, 38, '2023-07-26 20:52:12', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (39, '系统管理', 0, NULL, 39, '2023-07-26 20:52:19', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (40, '账户管理', 39, NULL, 40, '2023-07-26 20:52:27', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (41, '权限管理', 39, NULL, 41, '2023-07-26 20:52:41', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (42, '关于我们', 39, NULL, 42, '2023-07-26 20:52:51', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (43, '用户使用协议', 39, NULL, 43, '2023-07-26 20:53:06', NULL);
INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (44, '用户隐私政策', 39, NULL, 44, '2023-07-26 20:53:16', NULL);


*/
