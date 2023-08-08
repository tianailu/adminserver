package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/api/admin/user/services"
	"github.com/tianailu/adminserver/pkg/common"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (h *UserController) FindUserList(c echo.Context) error {
	var (
		req  = &models.UserSearchParam{}
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

	users, pageNum, pageSize, total, err := h.userService.Find(ctx, req)
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

func (h *UserController) FindUserDetail(c echo.Context) error {
	var (
		resp = common.Response{
			Status: 0,
			Msg:    "OK",
		}
		ctx = c.Request().Context()
	)

	userIdParam := c.Param("user_id")

	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return err
	}

	user, err := h.userService.FindUserDetail(ctx, userId)
	if err != nil {
		return err
	}

	resp.Data = user

	return c.JSON(http.StatusOK, resp)
}

func (h *UserController) AddUser(c echo.Context) error {
	var (
		req  = &models.UserDetail{}
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

	if err := h.userService.AddUser(ctx, req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *UserController) CreateUid(c echo.Context) error {
	var (
		resp = common.Response{
			Status: 0,
			Msg:    "OK",
		}
		ctx = c.Request().Context()
	)

	uid, err := h.userService.CreateUserId(ctx)
	if err != nil {
		return err
	}

	resp.Data = struct {
		Uid int64 `json:"uid,omitempty"`
	}{
		Uid: uid,
	}

	return c.JSON(http.StatusOK, resp)
}