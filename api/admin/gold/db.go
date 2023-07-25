package gold

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// 交易额明细
type GoldTradeList struct {
	ID          int                   `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
	UserID      int64                 `json:"user_id" gorm:"comment:用户id"`
	Amount      int64                 `json:"amount" gorm:"comment:用户此次交易的额度"`
	Type        int                   `json:"type" gorm:"comment:用户此次交易的种类"`
	TotalAmount int64                 `json:"total_amount" gorm:"comment:用户的总账户余额"`
	Extra       string                `json:"extra" gorm:"comment:额外信息"`
	CreatedTime time.Time             `json:"created_time" gorm:"index;type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	DeletedTime time.Time             `json:"deleted_time" gorm:"type:datetime;comment:删除时间"`
	IsDel       soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag;comment:删除标志，取值为[0:使用中, 1:已删除]"`
}

// 邀请交易额明细
type InviteInfo struct {
	ID            int                   `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
	InviteTradeID int64                 `json:"invite_trade_id"  gorm:"comment:邀请交易单ID"`
	HelpTradeID   int64                 `json:"help_trade_id" gorm:"comment:帮助交易单ID"`
	InviteCode    string                `json:"invite_code"  gorm:"index;comment:邀请码"`
	InviteType    int                   `json:"invite_type" gorm:"comment:邀请类型"`
	Extra         string                `json:"extra" gorm:"comment:额外信息"`
	CreatedTime   time.Time             `json:"created_time" gorm:"index;type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	DeletedTime   time.Time             `json:"deleted_time" gorm:"type:datetime;comment:删除时间"`
	IsDel         soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag;comment:删除标志，取值为[0:使用中, 1:已删除]"`
}

// 聊天设置
type ChatSetting struct {
	ID                 int       `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
	NoticeType         int       `json:"notice_type" gorm:"comment:0:仅好友消息提醒 1:统一接受提醒"`
	NoticeFrequency    int       `json:"notice_frequency" gorm:"comment:多少分钟提醒一次用户"`
	NoDisturbSwitch    bool      `json:"no_disturb_switch" gorm:"comment:免打扰开关"`
	NoDisturbStartTime time.Time `json:"no_disturb_start_time" gorm:"type:datetime;comment:免打扰开始时间"`
	NoDisturbEndTime   time.Time `json:"no_disturb_end_time" gorm:"type:datetime;comment:免打扰结束时间"`
	LoginNoticeText    string    `json:"login_notice_text" gorm:"comment:首次登陆提示"`
	FriendNoticeText   string    `json:"friend_notice_text" gorm:"comment:好友聊天框提示"`
	LetterNoticeText   string    `json:"letter_notice_text" gorm:"comment:飞鸽传书聊天框提示"`
	ShowHasReadType    int       `json:"show_has_read_type" gorm:"comment:0：消息显示已读 1：消息不显示已读"`
	AcceptNoticeType   int       `json:"accept_notice_type" gorm:"comment:0：所有人，1：仅好友，2：我关注的人"`
}

// 金币设置
type GoldSetting struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
	GoldName   int    `json:"gold_name" gorm:"size:9;not null;comment:金币名称"`
	InfoAmount int    `json:"info_amount" gorm:"default:200;comment:完成资料后给的金币奖励"`
	SignIn     string `json:"sign_in" gorm:"comment:签到相关设置"`
	Invite     string `json:"invite" gorm:"comment:邀请相关设置"`
}

// 金币设置
type RDBGoldSetting struct {
	ID         int           `json:"id" `
	GoldName   int           `json:"gold_name" `
	InfoAmount int           `json:"info_amount"`
	SignIn     SignInSetting `json:"sign_in" `
	Invite     InviteSetting `json:"invite" `
}
type SignInSetting struct {
	Switch       bool `json:"switch"`
	DailyGold    int  `json:"daily_gold"`
	ContinueDays int  `json:"continue_days"`
	ContinueGold int  `json:"continue_gold"`
	LoopSwitch   bool `json:"loop_switch"`
}

type InviteSetting struct {
	DailyLimit    int64 `json:"daily_limit"`
	TotalLimit    int64 `json:"total_limit"`
	InviteeAmount int64 `json:"invitee_amount"`
	InviterAmount int64 `json:"inviter_amount"`
}
