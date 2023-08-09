package models

type UserDetail struct {
	Id               uint    `json:"id,optional"`
	AccountId        string  `json:"account_id,optional"`
	UserId           int64   `json:"user_id,optional"`
	Name             string  `json:"name,optional"`
	Avatar           string  `json:"avatar,optional"`
	Gender           int8    `json:"gender,optional"`
	Birthday         int64   `json:"birthday,optional"`
	Constellation    string  `json:"constellation,optional"`
	Height           float32 `json:"height,optional"`
	Weight           float32 `json:"weight,optional"`
	Education        int8    `json:"education,optional"`
	EduStatus        int8    `json:"edu_status,optional"`
	School           string  `json:"school,optional"`
	Work             string  `json:"work,optional"`
	Company          string  `json:"company,optional"`
	Income           int8    `json:"income,optional"`
	Residence        string  `json:"residence,optional"`
	Hometown         string  `json:"hometown,optional"`
	MobilePhone      string  `json:"mobile_phone,optional"`
	IdentityTag      int8    `json:"identity_tag,optional"`
	IsVip            int8    `json:"is_vip,optional"`
	VipTag           int8    `json:"vip_tag,optional"`
	Recommend        int8    `json:"recommend,optional"`
	RegisterPlace    string  `json:"register_place,optional"`
	RegisterSource   int8    `json:"register_source,optional"`
	RegisterTime     int64   `json:"register_time,optional"`
	AuditStatus      int8    `json:"audit_status,optional"`
	UserStatus       int8    `json:"user_status,optional"`
	TotalUsageTime   int64   `json:"total_usage_time,optional"`
	Habit            string  `json:"habit,optional"`
	ConsumptionView  string  `json:"consumption_view,optional"`
	FamilyBackground string  `json:"family_background,optional"`
	Interest         string  `json:"interest,optional"`
	LoveView         string  `json:"love_view,optional"`
	TargetAppearance string  `json:"target_appearance,optional"`
	BeImpressed      string  `json:"be_impressed,optional"`
}

type UserSearchParam struct {
	Keywords        string `query:"keywords,optional"`          // 关键字，用户ID/昵称/用户名
	Gender          int8   `query:"gender,optional"`            // 性别，取值为[0:全部, 1:男, 2:女]
	IdentityTag     int8   `query:"identity_tag,optional"`      // 身份标签，取值为[0:全部, 1:母胎单身, 2:未婚单身, 3:离异无孩, 4:离异带孩, 5:离异不带孩, 6:丧偶]
	IsVip           int8   `query:"is_vip,optional"`            // 是否vip，取值为[0:全部, 1:是, 2:否]
	VipTag          int8   `query:"vip_tag,optional"`           // VIP标签
	AuditStatus     int8   `query:"audit_status,optional"`      // 基础信息审核状态，取值为[0:未知, 1:待审（首次申请审核）, 2: 再审核（非首次申请审核）, 3:通过, 4:不通过]
	Recommend       int8   `query:"recommend,optional"`         // 推荐，取值为[0:全部, 1:是, 2:否]
	RegisterPlace   string `query:"register_place,optional"`    // 注册地
	RegisterSource  int8   `query:"register_source,optional"`   // 注册来源，取值为[0:全部, 1:APP, 2:小程序, 3:群组, 4:二维码, 5:管理后台]
	RegisterStartAt int64  `query:"register_start_at,optional"` // 开始时间
	RegisterEndAt   int64  `query:"register_end_at,optional"`   // 结束时间
	PageNum         int    `query:"page_num,optional"`          // 第几页
	PageSize        int    `query:"page_size,optional"`         // 每页条数
}

type UserList struct {
}

type UserListItem struct {
	UserId         int64   `json:"user_id,optional"`
	Name           string  `json:"name,optional"`
	Gender         int8    `json:"gender,optional"`
	AuditStatus    int8    `json:"audit_status,optional"`
	IdentityTag    int8    `json:"identity_tag,optional"`
	IsVip          int8    `json:"is_vip,optional"`
	VipTag         int8    `json:"vip_tag,optional"`
	RechargeAmount float32 `json:"recharge_amount,optional"`
	RemainingCoins int64   `json:"remaining_coins,optional"`
	ConsumeCoins   int64   `json:"consume_coins,optional"`
	Recommend      int8    `json:"recommend,optional"`
	RegisterPlace  string  `json:"register_place,optional"`
	RegisterSource int8    `json:"register_source,optional"`
	RegisterTime   int64   `json:"register_time,optional"`
	DurationOfUse  int64   `json:"duration_of_use,optional"`
	// TODO 交友相关数据
}
