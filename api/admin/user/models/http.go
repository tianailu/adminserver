package models

type UserDetail struct {
	Id               uint    `json:"id,optional"`
	AccountId        string  `json:"account_id,optional"`
	Uid              int64   `json:"uid,optional"`
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
	Income           string  `json:"income,optional"`
	Residence        string  `json:"residence,optional"`
	Hometown         string  `json:"hometown,optional"`
	MobilePhone      string  `json:"mobile_phone,optional"`
	IdentityTag      int8    `json:"identity_tag,optional"`
	VipTag           int8    `json:"vipTag,optional"`
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
