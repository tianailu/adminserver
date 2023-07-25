package gold

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	TradeType_Recharge       = iota // 充值
	TradeType_Invite                // 邀请
	TradeType_SignIn                // 签到
	TradeType_RealCert              // 真人认证
	TradeType_AcademicCert          // 学历认证
	TradeType_Authentication        // 身份认证
	TradeType_FirstDate             // 用户首次发布周末约会
	TradeType_WeeklyDate            // 用户每周发布周末约会
	TradeType_FirstFriend           // 用户首次申请好友
	TradeType_WeeklyFriend          // 用户每周申请好友
	TradeType_AllInfo               // 用户完成所有资料填写
	TradeType_Heart                 // 发布心声
	TradeType_SameVoice             // 发布共鸣
	TradeType_UnLockLetter          // 解锁飞鸽传书

)

const (
	TradeDBName         = "tb_gold_trade"
	InviteDBName        = "tb_gold_invite"
	ChatSettingDBName   = "tb_chat_setting"
	InviteSettingDBName = "tb_invite_setting"
)

type SetInviteReq struct {
	InviteCode   string `json:"invite_code"`
	InviteUserID int64  `json:"invite_user_id"`
	HelpUserID   int64  `json:"help_user_id"`
	InviteType   int    `json:"invite_type"`
	jwt.RegisteredClaims
}

// 邀请记录数据
type GetInviteReq struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	InviteCode  string    `json:"invite_code"`
	InviterName string    `json:"inviter_name"`
	InviterGold int64     `json:"inviter_gold"`
	HelperName  string    `json:"helper_name"`
	HelperGold  int64     `json:"helper_gold"`
	time        time.Time `json:"time"`
}

// 邀请码记录数据
type GetInviteCodeReq struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	InviteCode  string    `json:"invite_code"`
	InviterName string    `json:"inviter_name"`
	InviteNum   int64     `json:"invite_num"`
	time        time.Time `json:"time"`
}

// 邀请码记录数据
type GetShareReq struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	InviteCode  string    `json:"invite_code"`
	InviterName string    `json:"inviter_name"`
	ShareType   int       `json:"share_type"`
	time        time.Time `json:"time"`
}

type SetTradeReq struct {
	UserID      int64 `json:"user_id" `
	Amount      int64 `json:"amount"`
	Type        int   `json:"type"`
	TotalAmount int64 `json:"total_amount"`
	jwt.RegisteredClaims
}

type SetSettingReq struct {
	Param string `json:"param" `
	Value string `json:"value"`
	jwt.RegisteredClaims
}

type GetTradeResp struct {
	Data interface{} `json:"data"`
	Page int64       `json:"page"`
}

type GetSettingResp struct {
	Data interface{} `json:"data"`
}
