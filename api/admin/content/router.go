package systemsetting

import (
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/tianailu/adminserver/api/admin/content/controller"
	"github.com/tianailu/adminserver/api/admin/content/domain"
	"github.com/tianailu/adminserver/pkg/db/mysql"
)

func init() {
	go initContentMgrData()
}

func initContentMgrData() error {
	time.Sleep(time.Second * 5)
	log.Printf("begin init content managemnt  table and data")

	tbTag := domain.Tag{}
	err := mysql.GetDb().AutoMigrate(&tbTag)
	if err != nil {
		log.Printf("create content managemnt table[%s] failed,err: %s", tbTag.TableName(), err)
		return err
	}

	tbGreet := domain.Greet{}
	err = mysql.GetDb().AutoMigrate(&tbGreet)
	if err != nil {
		log.Printf("create system setting table[%s] failed,err: %s", tbGreet.TableName(), err)
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
