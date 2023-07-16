package auth

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/config"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/utility/crypto"
	"github.com/tianailu/adminserver/pkg/utility/snowflake"
	"net/http"
)

func AddAdmin(c echo.Context) error {
	var (
		req = struct {
			Account  string `json:"account"`
			Password string `json:"password"`
			Name     string `json:"name"`
		}{}
		resp = common.Response{}
	)

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	account := &Account{
		UserId:      snowflake.GetNode().Generate().String(),
		Account:     req.Account,
		Password:    crypto.GetSha256String(req.Password, config.AuthConf["admin_password_salt"]),
		AccountType: "ADMIN",
		Name:        req.Name,
		Role:        "",
		Status:      0,
		DeletedAt:   sql.NullTime{Valid: false},
	}

	err := account.Create()
	if err != nil {
		c.Logger().Errorf("Failed to create admin account, error: %s", err.Error())
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
