package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/api/admin/user/services"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/errors"
	"net/http"
)

type VipController struct {
	vipService *services.VipService
}

func NewVipController() *VipController {
	return &VipController{
		vipService: services.NewVipService(),
	}
}

func (h *VipController) AddVipTag(c echo.Context) error {
	var (
		req  = &models.VipTagDetail{}
		resp = common.NewOkResponse()
		ctx  = c.Request().Context()
	)

	if err := c.Bind(req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}

	vipTagId, err := h.vipService.AddVipTag(ctx, req)
	if err != nil {
		return err
	}

	resp.Data = &struct {
		Id int32 `json:"id"`
	}{
		Id: vipTagId,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *VipController) BatchDeleteVipTag(c echo.Context) error {
	var (
		req  = &models.VipTagCommonParam{}
		resp = common.NewOkResponse()
		ctx  = c.Request().Context()
	)

	if err := c.Bind(req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	vipTagIds := req.VipTagIds
	if len(vipTagIds) <= 0 {
		return errors.ParamError
	}

	err := h.vipService.BatchDeleteVipTag(ctx, vipTagIds)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *VipController) UpdateVipTag(c echo.Context) error {
	var (
		req  = &models.VipTagDetail{}
		resp = common.NewOkResponse()
		ctx  = c.Request().Context()
	)

	if err := c.Bind(req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	if req.Id <= 0 || req.TotalRechargeAmount.LessThan(decimal.NewFromFloat(0.0)) {
		return errors.ParamError
	}

	err := h.vipService.UpdateVipTag(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *VipController) FindVipTagList(c echo.Context) error {
	var (
		req  = &common.SearchParam{}
		resp = common.NewOkResponse()
		ctx  = c.Request().Context()
	)

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("Bind req param error: %s", err.Error())
		return err
	}

	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	vipTags, pageNum, pageSize, total, err := h.vipService.Find(ctx, req)
	if err != nil {
		return err
	}

	resp.Data = common.PageData{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
		List:     common.ToAnySlice(vipTags),
	}

	return c.JSON(http.StatusOK, resp)
}
