package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/tianailu/adminserver/api/admin/systemsetting/domain"
	"github.com/tianailu/adminserver/api/admin/systemsetting/domain/req"
	"github.com/tianailu/adminserver/api/admin/systemsetting/service"
	"github.com/tianailu/adminserver/pkg/common"
)

func NewRoleController() *RoleController {
	return &RoleController{
		roleSvc: service.NewRoleService(),
		permSvc: service.NewPermissionService(),
	}
}

type RoleController struct {
	roleSvc service.RoleService
	permSvc service.PermissionService
}

// 添加角色  POST /system-setting/roles
func (rc *RoleController) SaveRole(c echo.Context) error {
	role := new(domain.Role)
	err := c.Bind(&role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	err = rc.roleSvc.SaveRole(role)
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
func (rc *RoleController) GetAllRoles(c echo.Context) error {
	var roles []domain.Role
	roles, err := rc.roleSvc.GetAllRoles()
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	return c.JSON(http.StatusOK, common.ResponseSuccessWithData(roles))
}

// 分页获取角色列表 /system-setting/roles/page
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

// 删除角色 /system-setting/roles/:id
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

// 获取角色权限详情 /system-setting/roles/:roleId/permissions
func (rc *RoleController) GetRolePermissions(c echo.Context) error {
	roleId := c.Param("roleId")
	intRoleId, err := strconv.Atoi(roleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	userId := "1"
	r, err := rc.roleSvc.GetUserRolePermissions(rc.permSvc, userId, intRoleId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("获取角色权限详情失败"))
	}
	retVal := &common.Response{
		Data: r,
	}
	return c.JSON(http.StatusOK, retVal)
}

// 保存角色权限 POST /system-setting/role/:roleId/permissions
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
	err = rc.roleSvc.SaveRolePermissions(rId, permissionIds)

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg(err.Error()))
	}

	return c.JSON(http.StatusOK, common.ResponseSuccess())
}
