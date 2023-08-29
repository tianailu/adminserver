package content_mgr

import (
	"github.com/labstack/echo/v4"
	"log"
	"time"

	"github.com/tianailu/adminserver/api/admin/content_mgr/controller"
	"github.com/tianailu/adminserver/api/admin/content_mgr/domain"
	"github.com/tianailu/adminserver/pkg/db/mysql"
)

func init() {
}

func InitTable() {
	go initContentMgrData()
}

func initContentMgrData() error {
	time.Sleep(time.Second * 5)
	log.Printf("begin init content_mgr managemnt  table and data")

	tbTag := domain.Tag{}
	err := mysql.GetDB().AutoMigrate(&tbTag)
	if err != nil {
		log.Printf("create content_mgr managemnt table[%s] failed,err: %s", tbTag.TableName(), err)
		return err
	}

	tbGreet := domain.Greet{}
	err = mysql.GetDB().AutoMigrate(&tbGreet)
	if err != nil {
		log.Printf("create content_mgr managemnt table[%s] failed,err: %s", tbGreet.TableName(), err)
		return err
	}

	return nil
}

func InitRouter(e *echo.Group) {
}

func InitGroupAdminRouter(e *echo.Group) {

	tagCtl := controller.NewTagController()
	e.POST("/content-mgr/tag", tagCtl.SaveTag)
	e.PUT("/content-mgr/tag/:id", tagCtl.UpdateTag)
	e.DELETE("/content-mgr/tag/:id", tagCtl.DeleteTagById)
	e.POST("/content-mgr/tag/delete", tagCtl.BatchDeleteTags)
	e.PUT("/content-mgr/tag/:id/enable", tagCtl.EnableTag)
	e.PUT("/content-mgr/tag/:id/disable", tagCtl.DisableTag)
	e.POST("/content-mgr/tag/query", tagCtl.GetTagsPage)

	greetCtl := controller.NewGreetController()
	e.POST("/content-mgr/greet", greetCtl.SaveGreet)
	e.PUT("/content-mgr/greet/:id", greetCtl.UpdateGreet)
	e.DELETE("/content-mgr/greet/:id", greetCtl.DeleteGreetById)
	e.POST("/content-mgr/greet/delete", greetCtl.BatchDeleteGreets)
	e.PUT("/content-mgr/greet/:id/enable", greetCtl.EnableGreet)
	e.PUT("/content-mgr/greet/:id/disable", greetCtl.DisableGreet)
	e.POST("/content-mgr/greet/query", greetCtl.GetGreetsPage)
}
