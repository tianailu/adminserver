package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/config"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/utility/crypto"
	"gorm.io/gorm"
	"net/http"
	"unicode/utf8"
)

func UpdateAdmin(c echo.Context) error {
	var (
		req = struct {
			MobilePhone string `json:"mobile_phone,optional"`
			Account     string `json:"account"`
			NewPassword string `json:"new_password,optional"`
			Name        string `json:"name,optional"`
			Avatar      string `json:"avatar,optional"`
			Status      int8   `json:"status,default=-1,optional"`
			Remark      string `json:"remark,optional"`
		}{}
		resp = common.Response{
			Status: 0,
			Msg:    "OK",
		}
	)

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	if len(req.Account) <= 0 {
		resp.Status, resp.Msg = 1, "参数 account 不能为空"
		return c.JSON(http.StatusOK, resp)
	}

	accountRepo := NewAccountRepo(mysql.GetDB(), c.Logger())
	account, err := accountRepo.FindByAccount(req.Account, DefaultAccountType)
	if err == gorm.ErrRecordNotFound {
		resp.Status, resp.Msg = 1, "指定账号不存在"
		return c.JSON(http.StatusOK, resp)
	} else if err != nil {
		c.Logger().Errorf("Failed to query account info, account: %s, type: %s, error: %s", req.Account, DefaultAccountType, err)
		resp.Status, resp.Msg = 1, "内部异常"
		return c.JSON(http.StatusOK, resp)
	}

	isUpdate := false
	if utf8.RuneCountInString(req.MobilePhone) > 0 {
		account.MobilePhone = req.MobilePhone
		isUpdate = true
	}
	if utf8.RuneCountInString(req.NewPassword) > 0 {
		account.Password = crypto.GetSha256String(req.NewPassword, config.AuthConf.AdminPasswordSalt)
		isUpdate = true
	}
	if utf8.RuneCountInString(req.Name) > 0 {
		account.Name = req.Name
		isUpdate = true
	}
	if utf8.RuneCountInString(req.Avatar) > 0 {
		account.Avatar = req.Avatar
		isUpdate = true
	}
	if req.Status >= 0 && req.Status <= 1 {
		account.Status = req.Status
		isUpdate = true
	}
	if utf8.RuneCountInString(req.Remark) > 0 {
		account.Remark = req.Remark
	}

	if isUpdate {
		if err := accountRepo.Update(account); err != nil {
			c.Logger().Errorf("Failed to update account info, account: %s, error: %s", account, err.Error())
			resp.Status, resp.Msg = 1, "信息更新失败"
			return c.JSON(http.StatusOK, resp)
		}
	}

	return c.JSON(http.StatusOK, resp)
}
