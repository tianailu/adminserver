package auth

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/config"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/utility/crypto"
	"github.com/tianailu/adminserver/pkg/utility/snowflake"
	"gorm.io/gorm"
	"net/http"
	"unicode/utf8"
)

func AddAdmin(c echo.Context) error {
	var (
		req = struct {
			MobilePhone string `json:"mobile_phone,optional"`
			Account     string `json:"account"`
			Password    string `json:"password"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar,optional"`
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

	if len(req.Account) <= 0 || len(req.Password) <= 0 {
		resp.Status, resp.Msg = 1, "账号or密码不能为空"
		return c.JSON(http.StatusOK, resp)
	}

	if utf8.RuneCountInString(req.Name) <= 0 {
		resp.Status, resp.Msg = 1, "昵称不能为空"
		return c.JSON(http.StatusOK, resp)
	}

	accountRepo := NewAccountRepo(mysql.GetDB(), c.Logger())
	account, err := accountRepo.FindByAccount(req.Account, DefaultAccountType)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.Logger().Errorf("Failed to query account info, account: %s, type: %s, error: %s", req.Account, DefaultAccountType, err)
		resp.Status, resp.Msg = 1, "内部异常"
		return c.JSON(http.StatusOK, resp)
	} else if account != nil {
		resp.Status, resp.Msg = 1, "账号已存在"
		return c.JSON(http.StatusOK, resp)
	}

	account = &Account{
		AccountId:   snowflake.GetNode().Generate().String(),
		MobilePhone: req.MobilePhone,
		Account:     req.Account,
		Password:    crypto.GetSha256String(req.Password, config.AuthConf.AdminPasswordSalt),
		AccountType: DefaultAccountType,
		Name:        req.Name,
		Avatar:      req.Avatar,
		Status:      0,
		DeletedAt:   sql.NullTime{Valid: false},
		Remark:      req.Remark,
	}
	if err := accountRepo.Create(account); err != nil {
		c.Logger().Errorf("Failed to create admin account, error: %s", err.Error())
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
