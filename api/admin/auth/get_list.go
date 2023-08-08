package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"net/http"
)

func GetAdminList(c echo.Context) error {
	var (
		req = struct {
			PageNum  int    `query:"page_num"`
			PageSize int    `query:"page_size"`
			Name     string `query:"name"`
		}{}
		resp common.Response
	)

	type respData struct {
		PageNum  int            `json:"page_num"`
		PageSize int            `json:"page_size"`
		Total    int            `json:"total"`
		List     []*AccountInfo `json:"list"`
	}

	if err := c.Bind(req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	pageNum := req.PageNum
	pageSize := req.PageSize

	if pageNum <= 0 {
		pageNum = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	accountRepo := NewAccountRepo(mysql.GetDB(), c.Logger())
	accounts, err := accountRepo.Find(req.Name, pageNum, pageSize)
	if err != nil {
		c.Logger().Errorf("Failed to find account list with, name: %s, pageNum: %d, pageSize: %d", req.Name, pageNum, pageSize)
		return err
	}

	totalAdmin, err := accountRepo.TotalAdmin()
	if err != nil {
		return err
	}

	result := make([]*AccountInfo, 0)
	for _, a := range accounts {
		lastLoginAt := int64(0)
		if a.LastLoginAt.Valid {
			lastLoginAt = a.LastLoginAt.Time.UnixNano() / 1e6
		}
		result = append(result, &AccountInfo{
			Id:          a.Id,
			AccountId:   a.AccountId,
			MobilePhone: a.MobilePhone,
			Account:     a.Account,
			AccountType: a.AccountType,
			Name:        a.Name,
			Avatar:      a.Avatar,
			Status:      a.Status,
			LoginCount:  a.LoginCount,
			LastLoginIp: a.LastLoginIp,
			LastLoginAt: lastLoginAt,
			CreatedAt:   a.CreatedAt.UnixNano() / 1e6,
			UpdatedAt:   a.UpdatedAt.UnixNano() / 1e6,
			Remark:      a.Remark,
		})
	}

	resp.Data = &respData{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    int(totalAdmin),
		List:     result,
	}

	return c.JSON(http.StatusOK, resp)
}
