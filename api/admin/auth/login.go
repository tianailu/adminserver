package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/config"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/utility/crypto"
	"gorm.io/gorm"
	"net/http"
	"time"
	"unicode/utf8"
)

type AdminJwtClaims struct {
	AccountType string `json:"account_type"`
	Roles       []int  `json:"roles"`
	Source      string `json:"source"`
	Status      int8   `json:"status"`
	jwt.RegisteredClaims
}

func AdminLogin(c echo.Context) error {
	var (
		req = struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		resp = common.Response{
			Status: 0,
			Msg:    "OK",
		}
	)

	type respData struct {
		AccessToken string `json:"access_token,omitempty"`
	}

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	if utf8.RuneCountInString(req.Username) <= 0 || utf8.RuneCountInString(req.Password) <= 0 {
		resp.Status = 1
		resp.Msg = "用户名或者密码错误"
		return c.JSON(http.StatusOK, resp)
	}

	accountRepo := NewAccountRepo(mysql.GetDB(), c.Logger())
	account, err := accountRepo.FindByAccount(req.Username, "ADMIN")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Status, resp.Msg = 1, "用户名或者密码错误"
		return c.JSON(http.StatusOK, resp)
	} else if err != nil {
		resp.Status, resp.Msg = 1, "内部异常"
		return c.JSON(http.StatusOK, resp)
	} else if account.Status != 0 {
		resp.Status, resp.Msg = 1, "登录失败，当前账号已被禁用"
		return c.JSON(http.StatusOK, resp)
	}

	password := crypto.GetSha256String(req.Password, config.AuthConf.AdminPasswordSalt)
	if password != account.Password {
		resp.Status, resp.Msg = 1, "用户名或者密码错误"
		return c.JSON(http.StatusOK, resp)
	}

	now := time.Now()
	var claims = AdminJwtClaims{
		AccountType: account.AccountType,
		Source:      "ADMIN",
		Status:      account.Status,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        account.AccountId,
			Issuer:    "tianailu",
			Subject:   account.Account,
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 72)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(config.AuthConf.AdminSecretKey))
	if err != nil {
		return err
	}

	lastLoginIp := c.Request().Header.Get("HTTP_X_FORWARDED_FOR")
	if len(lastLoginIp) <= 0 {
		lastLoginIp = c.Request().Header.Get("REMOTE_ADDR")
	}
	if err := accountRepo.Login(account.AccountId, lastLoginIp, account.LoginCount+1); err != nil {
		resp.Status, resp.Msg = 1, "内部服务异常"
		return c.JSON(http.StatusOK, resp)
	}

	resp.Data = respData{
		AccessToken: accessToken,
	}

	return c.JSON(http.StatusOK, resp)
}
