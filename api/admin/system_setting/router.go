package system_setting

import (
	"github.com/labstack/echo/v4"
	"log"
	"time"

	"github.com/tianailu/adminserver/api/admin/system_setting/controller"
	"github.com/tianailu/adminserver/api/admin/system_setting/domain"
	"github.com/tianailu/adminserver/pkg/db/mysql"
)

func init() {

}

func InitTable() {
	go initSystemSettingData()
}

func initSystemSettingData() error {
	time.Sleep(time.Second * 5)
	log.Printf("begin init system setting table and data")

	tbSoftwareSetting := domain.SoftwareSetting{}
	err := mysql.GetDB().AutoMigrate(&tbSoftwareSetting)
	if err != nil {
		log.Printf("create system setting table[%s] failed,err: %s", tbSoftwareSetting.TableName(), err)
		return err
	}

	tbRole := domain.Role{}
	err = mysql.GetDB().AutoMigrate(&tbRole)
	if err != nil {
		log.Printf("create system setting table[%s] failed,err: %s", tbRole.TableName(), err)
		return err
	}

	tbUserRole := domain.UserRole{}
	err = mysql.GetDB().AutoMigrate(&tbUserRole)
	if err != nil {
		log.Printf("create system setting table[%s] failed,err: %s", tbUserRole.TableName(), err)
		return err
	}

	tbPermission := domain.Permission{}
	err = mysql.GetDB().AutoMigrate(&tbPermission)
	if err != nil {
		log.Printf("create system setting table[%s] failed,err: %s", tbPermission.TableName(), err)
		return err
	}

	err = mysql.GetDB().Exec("INSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (1, '用户管理', 0, NULL, 1, '2023-07-26 20:45:33', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (2, '用户列表', 1, NULL, 2, '2023-07-26 20:45:33', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (3, '举报管理', 1, NULL, 3, '2023-07-26 20:46:53', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (4, '禁言管理', 1, NULL, 4, '2023-07-26 20:47:02', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (5, '封号管理', 1, NULL, 5, '2023-07-26 20:47:14', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (6, '冻结管理', 1, NULL, 6, '2023-07-26 20:47:22', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (7, '注销列表', 1, NULL, 7, '2023-07-26 20:47:34', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (8, '交友管理-用户列表', 1, NULL, 8, '2023-07-26 20:47:59', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (9, '心动列表', 1, NULL, 9, '2023-07-26 20:48:07', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (10, '标签管理', 1, NULL, 10, '2023-07-26 20:48:12', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (11, '打招呼管理', 1, NULL, 11, '2023-07-26 20:48:21', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (12, '用户详情', 1, NULL, 12, '2023-07-26 20:48:29', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (13, 'VIP等级', 1, NULL, 13, '2023-07-26 20:48:34', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (14, 'VIP列表', 1, NULL, 14, '2023-07-26 20:48:42', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (15, '交友设置', 1, NULL, 15, '2023-07-26 20:48:50', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (16, '交友主页设置', 1, NULL, 16, '2023-07-26 20:48:55', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (17, '推荐页设置', 1, NULL, 17, '2023-07-26 20:49:02', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (18, '资料编辑页设置', 1, NULL, 18, '2023-07-26 20:49:12', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (19, '用户体验设置', 1, NULL, 19, '2023-07-26 20:49:17', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (20, '消息设置', 1, NULL, 20, '2023-07-26 20:49:22', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (21, '收费设置', 1, NULL, 21, '2023-07-26 20:49:27', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (22, '内容管理', 0, NULL, 22, '2023-07-26 20:49:33', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (23, '发帖列表', 22, NULL, 23, '2023-07-26 20:49:42', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (24, '评论列表', 22, NULL, 24, '2023-07-26 20:49:52', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (25, '飞鸽传书', 22, NULL, 25, '2023-07-26 20:50:01', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (26, '共鸣', 22, NULL, 26, '2023-07-26 20:50:09', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (27, '聊天管理', 0, NULL, 27, '2023-07-26 20:50:16', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (28, '聊天设置', 27, NULL, 28, '2023-07-26 20:50:23', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (29, '金币管理', 0, NULL, 29, '2023-07-26 20:50:40', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (30, '金币记录', 29, NULL, 30, '2023-07-26 20:50:47', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (31, '金币银行', 29, NULL, 31, '2023-07-26 20:50:53', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (32, '资料记录', 29, NULL, 32, '2023-07-26 20:51:02', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (33, '签到记录', 29, NULL, 33, '2023-07-26 20:51:12', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (34, '邀请记录', 29, NULL, 34, '2023-07-26 20:51:19', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (35, '邀请记录码', 29, NULL, 35, '2023-07-26 20:51:30', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (36, '金币设置', 29, NULL, 36, '2023-07-26 20:51:37', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (37, '资料设置', 29, NULL, 37, '2023-07-26 20:51:47', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (38, '分享记录', 29, NULL, 38, '2023-07-26 20:52:12', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (39, '系统管理', 0, NULL, 39, '2023-07-26 20:52:19', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (40, '账户管理', 39, NULL, 40, '2023-07-26 20:52:27', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (41, '权限管理', 39, NULL, 41, '2023-07-26 20:52:41', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (42, '关于我们', 39, NULL, 42, '2023-07-26 20:52:51', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (43, '用户使用协议', 39, NULL, 43, '2023-07-26 20:53:06', NULL);\nINSERT INTO `tal`.`tb_permission`(`id`, `name`, `parent_id`, `route`, `sequence`, `create_at`, `create_account_id`) VALUES (44, '用户隐私政策', 39, NULL, 44, '2023-07-26 20:53:16', NULL);").Error
	if err != nil {
		log.Printf("init permission table[%s] data failed,err: %s", tbPermission.TableName(), err)
	}

	tbRolePermission := domain.RolePermission{}
	err = mysql.GetDB().AutoMigrate(&tbRolePermission)
	if err != nil {
		log.Printf("create system setting table[%s] failed,err: %s", tbRolePermission.TableName(), err)
		return err
	}

	return nil
}

func InitRouter(e *echo.Group) {
}

func InitGroupAdminRouter(e *echo.Group) {

	ctl := controller.NewSoftwareSettingController()
	// type支持 about-us、user-agreement，user-privacy-policy
	e.POST("/system-setting/:type", ctl.AddOrUpdateSoftwareSetting)
	e.GET("/system-setting/:type", ctl.GetSoftwareSetting)
	roleCtl := controller.NewRoleController()
	e.POST("/system-setting/roles", roleCtl.SaveRole)
	e.DELETE("/system-setting/roles/:id", roleCtl.DeleteRole)
	e.POST("/system-setting/roles/delete", roleCtl.DeleteRoles)
	e.GET("/system-setting/roles", roleCtl.GetAllRoles)
	e.GET("/system-setting/role/:roleId/permissions", roleCtl.GetRolePermissions)
	e.POST("/system-setting/roles/page", roleCtl.GetRolesPage)
	e.POST("/system-setting/role/:roleId/permissions", roleCtl.SaveRolePermissions)
	e.GET("/user/permissions", roleCtl.GetLoginUserRolesWithPermissions)

	// perCtl := controller.NewPermissionController()
	// e.GET("/system-setting/permissions", perCtl.GetUserPermisions)
}
