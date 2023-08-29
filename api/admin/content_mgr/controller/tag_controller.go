package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/content_mgr/domain/req"
	"github.com/tianailu/adminserver/api/admin/content_mgr/service"
	"github.com/tianailu/adminserver/api/admin/utils"
	"github.com/tianailu/adminserver/pkg/common"
	"net/http"
	"strconv"
)

type TagController struct {
	tagSvc service.TagService
}

func NewTagController() *TagController {
	return &TagController{
		tagSvc: service.NewTagService(),
	}
}

// POST /content-mgr/tag
func (tc *TagController) SaveTag(c echo.Context) error {
	tagDto := req.SaveOrUpdateTagDto{}
	err := c.Bind(&tagDto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	tagName := tagDto.Name
	if len(tagName) > 10 {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("标签名长度大于10"))
	}
	accountId := utils.GetLoginUserAccountId(c)
	err = tc.tagSvc.SaveTag(accountId, tagName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg(err.Error()))
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

// PUT /content-mgr/tag/:tagId
func (tc *TagController) UpdateTag(c echo.Context) error {
	tagDto := req.SaveOrUpdateTagDto{}
	err := c.Bind(&tagDto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	tagName := tagDto.Name
	if len(tagName) > 10 {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("标签名长度大于10"))
	}
	tagId := c.Param("id")
	accountId := utils.GetLoginUserAccountId(c)
	err = tc.tagSvc.UpdateTag(accountId, tagId, tagName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg(err.Error()))
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

// DELETE /content-mgr/tag/:tagId
func (tc *TagController) DeleteTagById(c echo.Context) error {
	tagId := c.Param("id")
	err := tc.tagSvc.DeleteTagById(tagId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("删除失败"))
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

// POST /content-mgr/tag/delete
func (tc *TagController) BatchDeleteTags(c echo.Context) error {
	var tagIds []int
	err := c.Bind(&tagIds)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequest())
	}
	err = tc.tagSvc.BatchDeleteTags(tagIds)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg(err.Error()))
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

// PUT /content-mgr/tag/{id}/enable
func (tc *TagController) EnableTag(c echo.Context) error {
	tagIdStr := c.Param("id")
	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("参数错误"))
	}
	err = tc.tagSvc.EnableTag(tagId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg(err.Error()))
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

// PUT /content-mgr/tag/{id}/disable
func (tc *TagController) DisableTag(c echo.Context) error {
	tagIdStr := c.Param("id")
	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("参数错误"))
	}
	err = tc.tagSvc.DisableTag(tagId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("关闭失败"))
	}
	return c.JSON(http.StatusOK, common.ResponseSuccess())
}

func (tc *TagController) GetTagsPage(c echo.Context) error {
	tagReq := req.TagQueryReq{}
	bindErr := c.Bind(&tagReq)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("参数错误"))
	}
	total, data, err := tc.tagSvc.GetTagsPage(tagReq)
	var dataResp = common.ResponseData{}
	iCount, err := strconv.Atoi(strconv.FormatInt(total, 10))
	dataResp.Total = iCount
	var pages = 0
	if iCount%tagReq.PageSize == 0 {
		pages = iCount / tagReq.PageSize
	} else {
		pages = (iCount / tagReq.PageNum) + 1
	}
	dataResp.Data = data
	dataResp.Pages = pages
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ResponseBadRequestWithMsg("关闭失败"))
	}
	return c.JSON(http.StatusOK, dataResp)
}
