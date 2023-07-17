package systemsetting

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/tianailu/adminserver/api/admin/systemsetting/domain"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
)

type PermissionResponse struct {
	Id       int                   `json:"id"`
	Name     string                `json:"name"`
	Route    string                `json:"route"`
	ParentId int                   `json:"parentId"`
	Order    int                   `json:"order"`
	Child    []*PermissionResponse `json:"child" gorm:"-"`
}

// 获取权限列表
// 根据用户角色获取权限，求权限的交集,返回给前端
func GetUserPermisions(c echo.Context) error {

	var permissions []PermissionResponse

	err := mysql.GetDb().
		Table(domain.Permission.TableName(domain.Permission{})).
		Order("`order` ASC").
		Find(&permissions).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	log.Println("permissions :", permissions)

	idPermisionnReferMap := make(map[int]*PermissionResponse, len(permissions))
	for idx, elem := range permissions {
		idPermisionnReferMap[elem.Id] = &permissions[idx]

	}
	var result []*PermissionResponse

	for i := 0; i < len(permissions); i++ {
		id := permissions[i].Id
		pId := permissions[i].ParentId
		if idPermisionnReferMap[id].ParentId == 0 {
			result = append(result, idPermisionnReferMap[id])
		} else {
			idPermisionnReferMap[pId].Child = append(idPermisionnReferMap[pId].Child, &permissions[i])
		}

	}

	// 获取角色权限
	
	return c.JSON(http.StatusOK, common.ResponseSuccessWithData(result))

}
