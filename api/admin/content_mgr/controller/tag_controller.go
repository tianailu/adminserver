package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/content_mgr/service"
)

type TagController struct {
	tagSvc service.TagService
}

func NewTagController() *TagController {
	return &TagController{
		tagSvc: service.NewTagService(),
	}
}

// POST /content_mgr-mgr/tags
func (tc *TagController) SaveTag(c echo.Context) error {
	return nil
}

// PUT /content_mgr-mgr/tags/:tagId
func (tc *TagController) UpdateTag(c echo.Context) error {
	return nil
}

// DELETE /content_mgr-mgr/tags/:tagId
func (tc *TagController) DeleteTagById(c echo.Context) error {
	return nil
}

// POST /content_mgr-mgr/tags/delete
func (tc *TagController) BatchDeleteTags(c echo.Context) error {
	return nil
}

// PUT /content_mgr-mgr/tags/delete
func (tc *TagController) Tags(c echo.Context) error {
	return nil
}
