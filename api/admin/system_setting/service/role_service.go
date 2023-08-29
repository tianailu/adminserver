package service

import (
	"errors"
	"log"
	"time"

	"github.com/tianailu/adminserver/api/admin/system_setting/domain"
	"github.com/tianailu/adminserver/api/admin/system_setting/domain/req"
	"github.com/tianailu/adminserver/api/admin/system_setting/domain/resp"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm"
)

type RoleService struct {
	db *gorm.DB
}

func NewRoleService() RoleService {
	return RoleService{
		db: mysql.GetDB(),
	}
}

func (rs *RoleService) SaveRole(userId string, role *domain.Role) error {
	var count int64
	err := rs.db.
		Model(&domain.Role{}).
		Select("name").
		Where("name = ?", role.Name).
		Count(&count).Error
	if err != nil || count > 0 {
		return errors.New("角色名称已存在，请勿重复创建")
	}
	role.CreateAccountId = userId
	role.CreateAt = time.Now().UnixMilli()
	err = rs.db.Save(role).Error
	return err
}

func (rs *RoleService) GetAllRoles() ([]domain.Role, error) {
	var roles []domain.Role
	err := rs.db.
		Model(&domain.Role{}).
		Select("id", "name", "create_at").
		Order("create_at DESC").
		Find(&roles).Error
	return roles, err
}

func (rs *RoleService) GetRolesPage(reqParam *req.RolePageRequest) (cnt int64, roles []domain.Role, err error) {
	if reqParam.PageNum == 0 {
		reqParam.PageNum = 1
	}
	if reqParam.PageSize == 0 {
		reqParam.PageSize = 10
	}
	offset := (reqParam.PageNum - 1) * reqParam.PageSize
	err = rs.db.
		Model(&domain.Role{}).
		Select("id", "name").
		Order("create_at DESC").
		Offset(offset).
		Limit(reqParam.PageSize).
		Find(&roles).Error
	if err != nil {
		log.Println("查询角色列表异常", err)
		return -1, nil, errors.New("查询角色列表异常")
	}
	var count int64
	err = rs.db.
		Model(&domain.Role{}).
		Select("id").
		Order("create_at DESC").
		Count(&count).Error
	if err != nil {
		log.Println("查询角色总数异常", err)
		return -1, nil, errors.New("查询角色总数异常")
	}
	return count, roles, err
}

func (rs *RoleService) DeleteRole(role domain.Role) error {
	e := rs.db.Transaction(func(tx *gorm.DB) error {
		// delete role
		err := rs.db.Delete(&role).Error
		if err != nil {
			return err
		}
		// delete user role
		err = rs.db.Where("role_id = ?", role.Id).Delete(&domain.UserRole{}).Error
		if err != nil {
			return err
		}
		// delete role permission
		err = rs.db.Where("role_id = ?", role.Id).Delete(&domain.RolePermission{}).Error
		if err != nil {
			return err
		}
		return nil
	})
	return e
}

func (rs *RoleService) DeleteRoles(roleIds []int) error {
	e := rs.db.Transaction(func(tx *gorm.DB) error {
		err := rs.db.Where("id in (?)", roleIds).Delete(&domain.Role{}).Error
		if err != nil {
			return err
		}
		err = rs.db.Where("role_id in (?)", roleIds).Delete(&domain.RolePermission{}).Error
		if err != nil {
			return err
		}
		err = rs.db.Where("role_id in (?)", roleIds).Delete(&domain.UserRole{}).Error
		return err
	})
	return e
}

func (rs *RoleService) SaveRolePermissions(userId string, roleId int, pemIds []int) error {

	e := rs.db.Transaction(func(tx *gorm.DB) error {
		role := domain.Role{}
		err := rs.db.
			Model(&domain.Role{}).
			Where("id = ?", roleId).
			First(&role).Error
		if err != nil {
			return errors.New("角色不存在")
		}
		err = rs.db.Transaction(func(tx *gorm.DB) error {
			var permissionCount int64
			err = rs.db.
				Model(&domain.Permission{}).
				Where("id in (?) and parent_id != 0", pemIds).
				Count(&permissionCount).Error
			if err != nil || int(permissionCount) != len(pemIds) {
				return errors.New("权限列表异常，包含不存在的权限")
			}
			err = rs.db.Where("role_id = ? ", roleId).Delete(&domain.RolePermission{}).Error
			if err != nil {
				log.Printf("删除角色%d权限异常", roleId)
				return err
			}
			var rolePermits []domain.RolePermission
			for _, permissionId := range pemIds {
				temp := domain.RolePermission{
					RoleId:          roleId,
					PermissionId:    permissionId,
					CreateAccountId: userId,
					CreateAt:        time.Now().UnixMilli(),
				}
				rolePermits = append(rolePermits, temp)
			}

			err = rs.db.
				Model(&domain.RolePermission{}).
				Save(rolePermits).Error
			return err
		})
		return err
	})

	return e
}

func (rs *RoleService) GetUserRolePermissions(permSvc PermissionService, roleId int) ([]*resp.RolePermissionDetail, error) {

	role := &domain.Role{}
	e := rs.db.Where("id = ?", roleId).First(role).Error
	if role.Id == 0 && e != nil {
		return nil, errors.New("角色不存在")
	}
	// 查询所有权限信息
	permAll, err := permSvc.GetAllPermissions()
	if err != nil {
		return nil, errors.New("获取权限列表失败")
	}
	var rolePerms []domain.RolePermission
	// 查询用户权限信息
	err = rs.db.Model(&domain.RolePermission{}).
		Where("role_id = ?", roleId).
		Order("permission_id ASC").
		Find(&rolePerms).Error
	if err != nil {
		return nil, err
	}
	// 返回权限信息
	var result []resp.RolePermissionDetail
	for _, v1 := range permAll {
		found := false
		for _, v2 := range rolePerms {
			if v1.Id == v2.PermissionId {
				rolePermDetail := &resp.RolePermissionDetail{
					Id:       v1.Id,
					Name:     v1.Name,
					Route:    v1.Route,
					ParentId: v1.ParentId,
					Child:    []*resp.RolePermissionDetail{},
					Sequence: v1.Sequence,
					Enable:   true,
				}
				result = append(result, *rolePermDetail)
				found = true
				break
			}

		}
		if !found {
			rolePermDetail := &resp.RolePermissionDetail{
				Id:       v1.Id,
				Name:     v1.Name,
				Route:    v1.Route,
				ParentId: v1.ParentId,
				Child:    []*resp.RolePermissionDetail{},
				Sequence: v1.Sequence,
			}

			if v1.ParentId == 0 {
				rolePermDetail.Enable = true
			} else {
				rolePermDetail.Enable = false
			}
			result = append(result, *rolePermDetail)
		}
	}
	retVal, err := permSvc.GetRolePermissionsTree(result)
	if err != nil {
		return nil, err
	}
	return retVal, err
}

func (rs *RoleService) GetUserFullyRolesAndPermissions(accountId string, permissionSvc PermissionService) ([]*resp.UserRolePermissions, error) {
	var roles []domain.Role
	roles, err := rs.getUserRoles(accountId)
	if err != nil {
		return nil, err
	}
	var userRolePermits []*resp.UserRolePermissions
	for _, v := range roles {
		permissions, err := rs.GetUserRolePermissions(permissionSvc, v.Id)
		if err != nil {
			return nil, err
		}
		item := &resp.UserRolePermissions{
			RoleId:       v.Id,
			RoleName:     v.Name,
			RoleCreateAt: v.CreateAt,
			Permissions:  permissions,
		}
		userRolePermits = append(userRolePermits, item)
	}
	return userRolePermits, err

}

func (rs *RoleService) getUserRoles(accountId string) ([]domain.Role, error) {
	var userRoles []domain.UserRole
	err := rs.db.Model(&domain.UserRole{}).Where("account_id = ?", accountId).Find(&userRoles).Error
	if err != nil {
		return nil, err
	}
	var roleIds = make([]int, len(userRoles))
	for _, role := range userRoles {
		roleIds = append(roleIds, role.RoleId)
	}
	var roles []domain.Role
	err = rs.db.Model(&domain.Role{}).Where("id in (?)", roleIds).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}
