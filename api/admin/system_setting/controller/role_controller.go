package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/tianailu/adminserver/api/admin/system_setting/domain"
	"github.com/tianailu/adminserver/api/admin/system_setting/domain/req"
	"github.com/tianailu/adminserver/api/admin/system_setting/service"
	"github.com/tianailu/adminserver/pkg/common"
)

type RoleController struct {
	roleSvc service.RoleService
	permSvc service.PermissionService
}

func NewRoleController() *RoleController {
	return &RoleController{
		roleSvc: service.NewRoleService(),
		permSvc: service.NewPermissionService(),
	}
}

// 添加角色  POST /system-setting/roles
func (rc *RoleController) SaveRole(c echo.Context) error {
	role := new(domain.Role)
	err := c.Bind(&role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	userId := utils.GetLoginUserAccountId(c)
	err = rc.roleSvc.SaveRole(userId, role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg(err.Error()))
	}

	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

type RoleList struct {
	Id   int
	Name string
}

// GetAllRoles 获取所有角色名称列表 /system-setting/roles
func (rc *RoleController) GetAllRoles(c echo.Context) error {
	var roles []domain.Role
	roles, err := rc.roleSvc.GetAllRoles()
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	return c.JSON(http.StatusOK, common.ResponseSuccessWithData(roles))
}

// GetRolesPage 分页获取角色列表 /system-setting/roles/page
func (rc *RoleController) GetRolesPage(c echo.Context) error {
	reqParam := new(req.RolePageRequest)
	err := c.Bind(reqParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	cnt, roles, err := rc.roleSvc.GetRolesPage(reqParam)
	iCount, err := strconv.Atoi(strconv.FormatInt(cnt, 10))
	if err != nil {
		log.Println("查询角色总数异常", err)
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	var dataResp = common.ResponseData{}
	dataResp.Data = roles
	dataResp.Total = iCount
	var pages int
	if iCount%reqParam.PageSize == 0 {
		pages = iCount / reqParam.PageSize
	} else {
		pages = iCount/reqParam.PageSize + 1
	}
	dataResp.Pages = pages
	return c.JSON(http.StatusOK, dataResp)
}

// DeleteRole 删除角色 /system-setting/roles/:id
func (rc *RoleController) DeleteRole(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	role := domain.Role{}
	role.Id = intId
	err = rc.roleSvc.DeleteRole(role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

// DeleteRoles 批量删除角色 POST /system-setting/roles/delete
func (rc *RoleController) DeleteRoles(c echo.Context) error {
	var roleIds []int
	err := c.Bind(&roleIds)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	err = rc.roleSvc.DeleteRoles(roleIds)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

// GetRolePermissions 获取角色权限详情 /system-setting/roles/:roleId/permissions
func (rc *RoleController) GetRolePermissions(c echo.Context) error {
	roleId := c.Param("roleId")
	intRoleId, err := strconv.Atoi(roleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	r, err := rc.roleSvc.GetUserRolePermissions(rc.permSvc, intRoleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg(err.Error()))
	}
	retVal := &common.Response{
		Data: r,
	}
	return c.JSON(http.StatusOK, retVal)
}

// SaveRolePermissions 保存角色权限 POST /system-setting/role/:roleId/permissions
func (rc *RoleController) SaveRolePermissions(c echo.Context) error {
	var permissionIds []int
	err := c.Bind(&permissionIds)
	roleId := c.Param("roleId")
	if len(permissionIds) == 0 {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("权限为空，请检查权限列表"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	rId, err := strconv.Atoi(roleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	userId := utils.GetLoginUserAccountId(c)
	err = rc.roleSvc.SaveRolePermissions(userId, rId, permissionIds)

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg(err.Error()))
	}

	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

// GetLoginUserRolesWithPermissions 保存角色权限 GET /user/permissions
func (rc *RoleController) GetLoginUserRolesWithPermissions(c echo.Context) error {
	userId := utils.GetLoginUserAccountId(c)
	result, e := rc.roleSvc.GetUserFullyRolesAndPermissions(userId, rc.permSvc)
	if e != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg(e.Error()))
	}
	return c.JSON(http.StatusOK, common.ResponseSuccessWithData(result))
}
