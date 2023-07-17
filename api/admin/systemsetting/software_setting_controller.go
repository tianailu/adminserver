package systemsetting

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/tianailu/adminserver/api/admin/systemsetting/domain"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm"
)

var setting_type = []string{"about-us", "user-agreement", "user-privacy-policy"}

// 添加或更新系统设置 about-us、user-agreement，user-privacy-policy
func AddOrUpdateSoftwareSetting(c echo.Context) error {
	// TODO 获取登录用户id
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(*jwt.MapClaims)

	settingType := c.Param("type")
	var (
		resp            common.ResponseNoData
		softwareSetting domain.SoftwareSetting
	)
	if !checkSettingType(settingType) {
		return c.JSON(http.StatusBadRequest, resp)
	}

	content := new(domain.SoftwareSettingRequest)
	if err := c.Bind(&content); err != nil {
		return err
	}

	softwareSetting.Content = content.Content
	softwareSetting.Type = settingType
	softwareSetting.CreateTime = int64(time.Now().UnixMilli())
	softwareSetting.UpdateTime = int64(time.Now().UnixMilli())

	err := mysql.GetDb().Transaction(func(tx *gorm.DB) error {
		return tx.Create(softwareSetting).Error
	})
	if err != nil {
		fmt.Println("保存失败")
	}
	resp.Status = 0
	resp.Msg = "ok"

	return c.JSON(http.StatusOK, resp)
}

// 查询系统设置 about-us、user-agreement，user-privacy-policy
func GetSoftwareSetting(c echo.Context) error {
	settingType := c.Param("type")

	if !checkSettingType(settingType) {
		return c.JSON(http.StatusBadRequest, nil)
	}
	var content string
	err := mysql.GetDb().
		Table(domain.SoftwareSetting.TableName(domain.SoftwareSetting{})).
		Select("content").
		Where("type", settingType).
		Find(&content).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{})
	}
	rd := common.Response{
		Status: 0,
		Data:   content,
	}
	return c.JSON(http.StatusOK, rd)
}

func checkSettingType(settingType string) (exist bool) {
	fmt.Printf("get type %s \n", settingType)
	var found = false
	for _, v := range setting_type {
		if v == settingType {
			found = true
			break
		}
	}
	return found
}
