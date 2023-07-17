package systemsetting

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/tianailu/adminserver/api/admin/systemsetting/domain"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
)

// 添加角色  POST /system-setting/roles
func SaveRole(c echo.Context) error {
	role := new(domain.Role)
	err := c.Bind(&role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	var count int64
	err = mysql.GetDb().Table(domain.Role.TableName(domain.Role{})).Select("name").Where("name = ?", role.Name).Count(&count).Error
	if err != nil || count > 0 {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("角色名称已存在，请勿重复创建"))
	}

	role.CreateTime = time.Now().UnixMilli()
	role.CreateUserId = 1
	err = mysql.GetDb().Save(role).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}

	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

type RoleList struct {
	Id   int
	Name string
}

// 获取所有角色名称列表 /system-setting/roles
func GetAllRoles(c echo.Context) error {
	var roles []RoleList
	err := mysql.GetDb().Table(domain.Role.TableName(domain.Role{})).Select("id", "name").Order("create_time DESC").Find(&roles).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	return c.JSON(http.StatusOK, common.ResponseSuccessWithData(roles))
}

type RolePageRequest struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
}

// 分页获取角色列表 /system-setting/roles/page
func GetRolesPage(c echo.Context) error {
	req := new(RolePageRequest)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	offset := (req.PageNum - 1) * req.PageSize
	var roles []RoleList
	err = mysql.GetDb().
		Table(domain.Role.TableName(domain.Role{})).
		Select("id", "name").
		Order("create_time DESC").
		Offset(offset).
		Limit(req.PageSize).
		Find(&roles).Error
	if err != nil {
		log.Println("查询角色列表异常", err)
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	var count int64
	err = mysql.GetDb().
		Table(domain.Role.TableName(domain.Role{})).
		Select("id").
		Order("create_time DESC").
		Count(&count).Error
	if err != nil {
		log.Println("查询角色总数异常", err)
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	var dataResp = common.ResponseData{}
	dataResp.Data = roles
	iCount, err := strconv.Atoi(strconv.FormatInt(count, 10))
	if err != nil {
		log.Println("查询角色总数异常", err)
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	dataResp.Total = iCount
	var pages int
	if iCount%req.PageSize == 0 {
		pages = iCount / req.PageSize
	} else {
		pages = iCount/req.PageSize + 1
	}
	dataResp.Pages = pages
	return c.JSON(http.StatusOK, dataResp)
}

// 删除角色 /system-setting/roles/:id
func DeleteRole(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	role := domain.Role{}
	role.Id = intId
	err = mysql.GetDb().Delete(&role).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

// 获取角色权限详情
func GetRolePermissions(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

type saveRolePermissionRequest struct {
	RoleId        int   `json:"roleId"`
	PermissionIds []int `json:"permissionIds"`
}

// 保存角色权限 POST /system-setting/role-permissions
func SaveRolePermissions(c echo.Context) error {
	req := saveRolePermissionRequest{}
	err := c.Bind(&req)
	if len(req.PermissionIds) == 0 {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("权限为空，请检查权限列表"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	role := domain.Role{}
	err = mysql.GetDb().
		Table(domain.Role.TableName(domain.Role{})).
		Where("id = ?", req.RoleId).
		First(&role).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("角色不存在"))
	}
	var permissionCount int64
	err = mysql.GetDb().
		Table(domain.Permission.TableName(domain.Permission{})).
		Where("id in (?) and parent_id != 0", req.PermissionIds).
		Count(&permissionCount).Error
	if err != nil || int(permissionCount) != len(req.PermissionIds) {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("权限列表异常，包含不存在的权限"))
	}

	var rolePermis []domain.RolePermission
	for _, permissionId := range req.PermissionIds {
		temp := domain.RolePermission{
			RoleId:       req.RoleId,
			PermissionId: permissionId,
			CreateTime:   time.Now().UnixMilli(),
			CreateUserId: 1, // fixme 登录后修复
		}
		rolePermis = append(rolePermis, temp)
	}

	err = mysql.GetDb().
		Table(domain.RolePermission.TableName(domain.RolePermission{})).
		Save(rolePermis).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.ResponseCommonFailed())
	}

	return c.JSON(http.StatusOK, nil)
}
