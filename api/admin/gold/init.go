package gold

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/pkg/db/mysql"
)

func initDB() {
	mysql.GetDB().Table(TradeDBName).AutoMigrate(&GoldTradeList{})
	mysql.GetDB().Table(InviteDBName).AutoMigrate(&InviteInfo{})
	mysql.GetDB().Table(ChatSettingDBName).AutoMigrate(&ChatSetting{})
}
func InitRouter(g *echo.Group) {
	initDB()
	g.POST("/gold/set", SetTrade)
	g.GET("/gold/get", GetTrade)
	g.GET("/chat/setting/get", GetChatSetting)
	g.POST("/chat/setting/set", SetChatSetting)
	g.GET("/gold/setting/get", GetGoldSetting)
	g.POST("/gold/setting/set", SetGoldSetting)
	g.POST("/gold/invite/set", SetInvite)
	g.GET("/gold/excel/invite", GetInviteExcel)
	g.GET("/gold/excel/invite_code", GetInviteCodeExcel)
	g.GET("/gold/excel/share", GetShareExcel)
	g.GET("/gold/excel/sign_in", GetSignInExcel)
}
