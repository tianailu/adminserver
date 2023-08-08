package controller

import (
	"github.com/labstack/echo"
	"github.com/tianailu/adminserver/api/admin/content/service"
)

type TagController struct {
	tagSvc service.TagService
}

func NewTagController() *TagController {
	return &TagController{
		tagSvc: service.NewTagService(),
	}
}


// POST /content-mgr/tags
func(tc *TagController) SaveTag(c echo.Context) error {
	return nil
}


// PUT /content-mgr/tags/:tagId
func(tc *TagController) UpdateTag(c echo.Context) error {
	return nil
}

// DELETE /content-mgr/tags/:tagId
func(tc *TagController) DeleteTagById(c echo.Context) error {
	return nil
}
// POST /content-mgr/tags/delete
func(tc *TagController) BatchDeleteTags(c echo.Context) error {
	return nil
}

// PUT /content-mgr/tags/delete
func(tc *TagController) Tags(c echo.Context) error {
	return nil
}