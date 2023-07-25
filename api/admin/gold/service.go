package gold

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm/clause"
)

func updateInviteInfo(c echo.Context, req *SetInviteReq) error {
	// 开始事务
	tx := mysql.GetDB().Begin()
	// 查询邀请的交易额
	var inviteInfo InviteSetting
	if err := tx.Table(InviteSettingDBName).First(&inviteInfo).Error; err != nil {
		tx.Rollback()
		c.Logger().Errorf("[updateInviteInfo] query MySQL key fail,err:%v", err)
		return err
	}

	// 向交易中插入两条记录，分别是邀请用户和帮助用户的账户变动
	// 开启读锁
	var inviterInfo GoldTradeList // 请替换为您的数据模型
	if err := tx.Where("user_id = ?", req.InviteUserID).Clauses(clause.Locking{Strength: "UPDATE"}).Last(&inviterInfo).Error; err != nil {
		tx.Rollback()
		c.Logger().Errorf("Failed to read lock last record: %v", err)
		return err
	}
	inviterInfo.Amount = inviteInfo.InviterAmount
	inviterInfo.TotalAmount += inviterInfo.Amount
	inviterInfo.Type = TradeType_Invite
	inviterResult := tx.Table(TradeDBName).Create(&inviterInfo)
	if inviterResult.Error != nil || inviterResult.RowsAffected == 0 {
		tx.Rollback()
		c.Logger().Errorf("Failed to insert invite user account change: %v", inviterResult.Error)
		return inviterResult.Error
	}
	var inviteeInfo GoldTradeList
	if err := tx.Where("user_id = ?", req.HelpUserID).Clauses(clause.Locking{Strength: "UPDATE"}).Last(&inviterInfo).Error; err != nil {
		tx.Rollback()
		c.Logger().Errorf("Failed to read lock last record: %v", err)
		return err
	}
	inviteeInfo.Amount = inviteInfo.InviterAmount
	inviteeInfo.TotalAmount += inviterInfo.Amount
	inviteeInfo.Type = TradeType_Invite
	InviteeResult := tx.Table(TradeDBName).Create(&inviteeInfo)
	if InviteeResult.Error != nil || InviteeResult.RowsAffected == 0 {
		tx.Rollback()
		c.Logger().Errorf("Failed to insert invite user account change: %v", InviteeResult.Error)
		return InviteeResult.Error
	}

	InviteResult := tx.Table(InviteDBName).Create(&InviteInfo{
		InviteTradeID: inviterResult.RowsAffected,
		HelpTradeID:   InviteeResult.RowsAffected,
		InviteCode:    req.InviteCode,
		InviteType:    req.InviteType,
		Extra:         "",
	})
	if InviteResult.Error != nil || InviteResult.RowsAffected == 0 {
		tx.Rollback()
		c.Logger().Errorf("Failed to insert invite user account change: %v", InviteResult.Error)
		return InviteResult.Error
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		c.Logger().Errorf("Failed to commit transaction: %v", err)
	}
	return nil
}
