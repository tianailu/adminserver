package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/api/admin/user/services"
	"github.com/tianailu/adminserver/pkg/common"
	"net/http"
)

type FriendController struct {
	friendService *services.FriendService
}

func NewFriendController() *FriendController {
	return &FriendController{
		friendService: services.NewFriendService(),
	}
}

func (h *FriendController) FindFriendshipList(c echo.Context) error {
	var (
		req  = &models.FriendSearchParam{}
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

	users, pageNum, pageSize, total, err := h.friendService.Find(ctx, req)
	if err != nil {
		return err
	}

	resp.Data = common.PageData{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
		List:     common.ToAnySlice(users),
	}

	return c.JSON(http.StatusOK, resp)
}
