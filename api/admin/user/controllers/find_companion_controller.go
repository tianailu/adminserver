package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/api/admin/user/services"
	"github.com/tianailu/adminserver/pkg/common"
	"net/http"
)

type FindCompanionController struct {
	findCompanionService *services.FindCompanionService
}

func NewFindCompanionController() *FindCompanionController {
	return &FindCompanionController{
		findCompanionService: services.NewFindCompanionService(),
	}
}

func (h *FindCompanionController) FindCompanionList(c echo.Context) error {
	var (
		req  = &models.FindCompanionSearchParam{}
		resp = common.Response{
			Status: 0,
			Msg:    "OK",
		}
		ctx = c.Request().Context()
	)

	if err := c.Bind(req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	findCompanions, pageNum, pageSize, total, err := h.findCompanionService.Find(ctx, req)
	if err != nil {
		return err
	}

	resp.Data = common.PageData{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
		List:     common.ToAnySlice(findCompanions),
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *FindCompanionController) FindCompanionTypeList(c echo.Context) error {
	var (
		req  = &models.CompanionTypeSearchParam{}
		resp = common.Response{
			Status: 0,
			Msg:    "OK",
		}
		ctx = c.Request().Context()
	)

	if err := c.Bind(req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	companionTypes, err := h.findCompanionService.FindCompanionTypeList(ctx, req)
	if err != nil {
		return err
	}

	resp.Data = &struct {
		List []any `json:"list"`
	}{
		List: common.ToAnySlice(companionTypes),
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *FindCompanionController) AddCompanionType(c echo.Context) error {
	var (
		req = &struct {
			CompanionTag      int8   `json:"companion_tag"`
			CompanionTypeName string `json:"companion_type_name"`
			Status            int8   `json:"status"`
		}{}
		resp = common.Response{
			Status: 0,
			Msg:    "OK",
		}
		ctx = c.Request().Context()
	)

	if err := c.Bind(req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	err := h.findCompanionService.AddCompanionType(ctx, req.CompanionTag, req.CompanionTypeName, req.Status)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *FindCompanionController) UpdateCompanionType(c echo.Context) error {
	var (
		req = &struct {
			CompanionTypes []*models.CompanionTypeListItem `json:"companionTypes"`
		}{}
		resp = common.Response{
			Status: 0,
			Msg:    "OK",
		}
		ctx = c.Request().Context()
	)

	if err := c.Bind(req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	err := h.findCompanionService.UpdateCompanionTypeList(ctx, req.CompanionTypes)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
