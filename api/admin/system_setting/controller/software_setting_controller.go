package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/utils"
	"net/http"

	"github.com/tianailu/adminserver/api/admin/system_setting/domain"
	"github.com/tianailu/adminserver/api/admin/system_setting/service"
	"github.com/tianailu/adminserver/pkg/common"
)

func NewSoftwareSettingController() *SoftwareSettingController {
	return &SoftwareSettingController{
		srv: service.NewSoftwareSettingService(),
	}
}

type SoftwareSettingController struct {
	srv service.SoftwareSettingService
}

var setting_type = []string{"about-us", "user-agreement", "user-privacy-policy"}

// 添加或更新系统设置 about-us、user-agreement，user-privacy-policy
func (sfc *SoftwareSettingController) AddOrUpdateSoftwareSetting(c echo.Context) error {

	loginUserId := utils.GetLoginUserAccountId(c)
	settingType := c.Param("type")
	var softwareSetting domain.SoftwareSetting

	if !checkSettingType(settingType) {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}

	content := new(domain.SoftwareSettingRequest)
	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}

	softwareSetting.Content = content.Content
	softwareSetting.Type = settingType

	err := sfc.srv.AddOrUpdateSoftwareSetting(loginUserId, &softwareSetting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

// 查询系统设置 about-us、user-agreement，user-privacy-policy
func (sfc *SoftwareSettingController) GetSoftwareSetting(c echo.Context) error {
	settingType := c.Param("type")

	if !checkSettingType(settingType) {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}

	content, err := sfc.srv.GetSoftwareSetting(settingType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.ResponseSuccess())
	}
	rd := common.Response{
		Status: 0,
		Data:   content,
	}
	return c.JSON(http.StatusOK, rd)
}

func checkSettingType(settingType string) (exist bool) {
	var found = false
	for _, v := range setting_type {
		if v == settingType {
			found = true
			break
		}
	}
	return found
}
