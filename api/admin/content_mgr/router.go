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

	ctl := controller.NewTagController()
	e.POST("/content-mgr/tags", ctl.SaveTag)
	e.POST("/content-mgr/tags", ctl.SaveTag)
}
