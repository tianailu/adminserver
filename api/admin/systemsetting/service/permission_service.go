package service

import (
	"errors"

	"github.com/tianailu/adminserver/api/admin/systemsetting/domain"
	"github.com/tianailu/adminserver/api/admin/systemsetting/domain/resp"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm"
)

type PermissionService struct {
	db *gorm.DB
}

func NewPermissionService() PermissionService {
	return PermissionService{
		db: mysql.GetDb(),
	}
}

func (ps *PermissionService) GetAllPermissions() (permissions []*resp.PermissionResponse, err error) {
	var permAll []*resp.PermissionResponse
	err = ps.db.
		Model(&domain.Permission{}).
		Order("`order` ASC").
		Find(&permAll).Error
	if err != nil {
		return nil, errors.New("获取权限列表失败")
	}
	return permAll, err
}

func (ps *PermissionService) GetAllPermissionsTree() (permissions []*resp.PermissionResponse, err error) {
	var permAll []resp.PermissionResponse

	e := ps.db.
		Model(&domain.Permission{}).
		Order("`order` ASC").
		Find(&permAll).Error
	if e != nil {
		return nil, errors.New("获取权限列表失败")
	}

	permAll = append(permAll)
	idPermisionnReferMap := make(map[int]*resp.PermissionResponse, len(permAll))
	for idx, elem := range permAll {
		idPermisionnReferMap[elem.Id] = &permAll[idx]

	}
	var result []*resp.PermissionResponse

	for i := 0; i < len(permAll); i++ {
		id := permAll[i].Id
		pId := permAll[i].ParentId
		if idPermisionnReferMap[id].ParentId == 0 {
			result = append(result, idPermisionnReferMap[id])
		} else {
			idPermisionnReferMap[pId].Child = append(idPermisionnReferMap[pId].Child, &permAll[i])
		}

	}
	return result, nil
}

func (ps *PermissionService) GetRolePermissionsTree(permissions []resp.RolePermissionDetail) (r []*resp.RolePermissionDetail,err error) {

	idPermisionnReferMap := make(map[int]*resp.RolePermissionDetail, len(permissions))
	for idx, elem := range permissions {
		idPermisionnReferMap[elem.Id] = &permissions[idx]

	}
	var result []*resp.RolePermissionDetail

	for i := 0; i < len(permissions); i++ {
		id := permissions[i].Id
		pId := permissions[i].ParentId
		if idPermisionnReferMap[id].ParentId == 0 {
			result = append(result, idPermisionnReferMap[id])
		} else {
			idPermisionnReferMap[pId].Child = append(idPermisionnReferMap[pId].Child, &permissions[i])
		}

	}
	return result,nil
}
