package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/api/admin/user/services"
	"github.com/tianailu/adminserver/pkg/common"
	"net/http"
)

type HeartbeatController struct {
	heartbeatService *services.HeartbeatService
}

func NewHeartbeatController() *HeartbeatController {
	return &HeartbeatController{
		heartbeatService: services.NewHeartbeatService(),
	}
}

func (h *HeartbeatController) FindHeartbeatList(c echo.Context) error {
	var (
		req  = &models.HeartbeatSearchParam{}
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

	heartbeats, pageNum, pageSize, total, err := h.heartbeatService.Find(ctx, req)
	if err != nil {
		return err
	}

	resp.Data = common.PageData{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
		List:     common.ToAnySlice(heartbeats),
	}

	return c.JSON(http.StatusOK, resp)
}
