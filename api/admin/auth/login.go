package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/config"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/utility/crypto"
	"gorm.io/gorm"
	"net/http"
	"time"
	"unicode/utf8"
)

var table = "tb_account"

type JwtCustomClaims struct {
	UserId      string `json:"user_id"`
	AccountType string `json:"account_type"`
	Role        string `json:"role"`
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
		resp common.Response
	)

	type respData struct {
		AccessToken string `json:"access_token,omitempty"`
	}

	c.Bind(&req)

	if utf8.RuneCountInString(req.Username) <= 0 || utf8.RuneCountInString(req.Password) <= 0 {
		resp.Status = 1
		resp.Msg = "用户名或者密码错误"
		return c.JSON(http.StatusOK, resp)
	}

	var account *Account
	result, err := account.FindByAccount(req.Username, "ADMIN")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Status = 1
		resp.Msg = "用户名或者密码错误"
		return c.JSON(http.StatusOK, resp)
	} else if err != nil {
		return err
	}

	password := crypto.GetSha256String(req.Password, config.AuthConf["admin_password_salt"])
	if password != result.Password {
		resp.Status = 1
		resp.Msg = "用户名或者密码错误"
		return c.JSON(http.StatusOK, resp)
	}

	var claims = &JwtCustomClaims{
		UserId:      result.UserId,
		AccountType: result.AccountType,
		Role:        result.Role,
		Source:      "ADMIN",
		Status:      result.Status,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   result.Account,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(config.AuthConf["admin_secret_key"]))
	if err != nil {
		return err
	}

	if err := account.Login(result.UserId); err != nil {
		resp.Status = 1
		resp.Msg = "服务异常"
		return c.JSON(http.StatusOK, resp)
	}

	resp.Status = 0
	resp.Msg = "OK"
	resp.Data = respData{
		AccessToken: accessToken,
	}

	return c.JSON(http.StatusOK, resp)
}
