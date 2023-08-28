package models

import (
	"database/sql"
)

type UserDetail struct {
	Id               int64   `json:"id,optional"`
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
	Work             int     `json:"work,optional"`
	Company          string  `json:"company,optional"`
	Income           int8    `json:"income,optional"`
	Residence        string  `json:"residence,optional"`
	Hometown         string  `json:"hometown,optional"`
	MobilePhone      string  `json:"mobile_phone,optional"`
	IdentityTag      int8    `json:"identity_tag,optional"`
	IsVip            int8    `json:"is_vip,optional"`
	VipTag           int32   `json:"vip_tag,optional"`
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

type FriendDetail struct {
	FriendRequestId     int64         `json:"friend_request_id"`
	SenderUserId        int64         `json:"sender_user_id"`
	SenderName          string        `json:"sender_name"`
	SenderAvatar        string        `json:"sender_avatar"`
	SenderGender        int8          `json:"sender_gender"`
	SenderIdentityTag   int8          `json:"sender_identity_tag"`
	ReceiverUserId      int64         `json:"receiver_user_id"`
	ReceiverName        string        `json:"receiver_name"`
	ReceiverAvatar      string        `json:"receiver_avatar"`
	ReceiverGender      int8          `json:"receiver_gender"`
	ReceiverIdentityTag int8          `json:"receiver_identity_tag"`
	MatchingStatus      int8          `json:"matching_status"`
	ReceiverConfirmTime sql.NullInt64 `json:"receiver_confirm_time"`
	RequestTime         int64         `json:"request_time"`
}

type HeartbeatDetail struct {
	HeartbeatRequestId  int64         `json:"heartbeat_request_id"`
	SenderUserId        int64         `json:"sender_user_id"`
	SenderName          string        `json:"sender_name"`
	SenderAvatar        string        `json:"sender_avatar"`
	SenderGender        int8          `json:"sender_gender"`
	SenderIdentityTag   int8          `json:"sender_identity_tag"`
	ReceiverUserId      int64         `json:"receiver_user_id"`
	ReceiverName        string        `json:"receiver_name"`
	ReceiverAvatar      string        `json:"receiver_avatar"`
	ReceiverGender      int8          `json:"receiver_gender"`
	ReceiverIdentityTag int8          `json:"receiver_identity_tag"`
	MatchingStatus      int8          `json:"matching_status"`
	ReceiverConfirmTime sql.NullInt64 `json:"receiver_confirm_time"`
	RequestTime         int64         `json:"request_time"`
}

type FindCompanionDetail struct {
	FindCompanionActivityId int64         `json:"find_companion_activity_id"`
	FindCompanionRequestId  uint          `json:"find_companion_request_id"`
	SenderUserId            int64         `json:"sender_user_id"`
	SenderName              string        `json:"sender_name"`
	SenderAvatar            string        `json:"sender_avatar"`
	SenderGender            int8          `json:"sender_gender"`
	SenderIdentityTag       int8          `json:"sender_identity_tag"`
	ReceiverUserId          int64         `json:"receiver_user_id"`
	ReceiverName            string        `json:"receiver_name"`
	ReceiverAvatar          string        `json:"receiver_avatar"`
	ReceiverGender          int8          `json:"receiver_gender"`
	ReceiverIdentityTag     int8          `json:"receiver_identity_tag"`
	CompanionTypeId         int           `json:"companion_type_id"`
	CompanionTypeName       string        `json:"companion_type_name"`
	MatchingStatus          int8          `json:"matching_status"`
	PublishingActivityTime  int64         `json:"publishing_activity_time"`
	RequestTime             sql.NullInt64 `json:"request_time"`
	ActivityCancelTime      sql.NullInt64 `json:"activity_cancel_time"`
	SenderConfirmTime       sql.NullInt64 `json:"sender_confirm_time"`
}

type UserFollowStat struct {
	UserId      int64 `json:"user_id"`
	FollowCount int64 `json:"follow_count"`
}

type UserFansStat struct {
	UserId    int64 `json:"user_id"`
	FansCount int64 `json:"fans_count"`
}

type UserFollowFansStat struct {
	UserId      int64 `json:"user_id"`
	FollowCount int64 `json:"follow_count"`
	FansCount   int64 `json:"fans_count"`
}

type FriendRequestStat struct {
	UserId              int64 `json:"user_id"`
	RequestCount        int64 `json:"request_count"`
	RequestSuccessCount int64 `json:"request_success_count"`
}

type HeartbeatRequestStat struct {
	UserId              int64 `json:"user_id"`
	RequestCount        int64 `json:"request_count"`
	RequestSuccessCount int64 `json:"request_success_count"`
}
