package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/tianailu/adminserver/api/admin/content_mgr/domain"
	"github.com/tianailu/adminserver/api/admin/content_mgr/domain/req"
	"github.com/tianailu/adminserver/api/admin/content_mgr/service"
	"github.com/tianailu/adminserver/pkg/common"
)

type GreetController struct {
	greetSvc *service.GreetService
}

func NewGreetController() *GreetController {
	return &GreetController{
		greetSvc: service.NewGreetService(),
	}
}

/*
add greet
POST /content_mgr-mgr/greets
*/
func (gctl *GreetController) SaveGreet(c echo.Context) error {
	var reqParam req.SaveOrUpdateGreetDto
	err := c.Bind(&reqParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	if len(reqParam.Content) == 0 {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("greet content_mgr is required"))
	}
	greet := domain.Greet{
		Content: reqParam.Content,
	}
	err = gctl.greetSvc.SaveOrUpdateGreet(greet)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

/*
update greet
PUT /content_mgr-mgr/greets/{id}
*/
func (gctl *GreetController) UpdateGreet(c echo.Context) error {
	var reqParam req.SaveOrUpdateGreetDto
	strId := c.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	err = c.Bind(&reqParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	if len(reqParam.Content) == 0 {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("greet content_mgr is required"))
	}
	greet := domain.Greet{
		Content: reqParam.Content,
		Id:      intId,
	}
	err = gctl.greetSvc.SaveOrUpdateGreet(greet)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

/*
delete greet by id
DELETE /content_mgr-mgr/greets/{id}
*/
func (gctl *GreetController) DeleteGreetById(c echo.Context) error {
	return nil
}

/*
batch delete greets
POST /content_mgr-mgr/greets/delete
*/
func (gctl *GreetController) BatchDeleteGreets(c echo.Context) error {
	return nil
}

/*
enable greet
PUT /content_mgr-mgr/greets/{id}/enable
*/
func (gctl *GreetController) EnableGreet(c echo.Context) error {
	return nil
}

/*
disable greet
PUT /content_mgr-mgr/greets/{id}/disable
*/
func (gctl *GreetController) DisableGreet(c echo.Context) error {
	return nil
}

/*
pagination get greets list, search by keyword
POST /content_mgr-mgr/greets
*/
func (gctl *GreetController) GetGreetsPage(c echo.Context) error {
	return nil
}
