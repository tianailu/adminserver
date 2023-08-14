package models

import (
	"database/sql"
)

type FriendDetail struct {
	FriendRequestId     uint         `json:"friend_request_id"`
	SenderUserId        int64        `json:"sender_user_id"`
	SenderName          string       `json:"sender_name"`
	SenderAvatar        string       `json:"sender_avatar"`
	SenderGender        int8         `json:"sender_gender"`
	SenderIdentityTag   int8         `json:"sender_identity_tag"`
	ReceiverUserId      int64        `json:"receiver_user_id"`
	ReceiverName        string       `json:"receiver_name"`
	ReceiverAvatar      string       `json:"receiver_avatar"`
	ReceiverGender      int8         `json:"receiver_gender"`
	ReceiverIdentityTag int8         `json:"receiver_identity_tag"`
	MatchingStatus      int8         `json:"matching_status"`
	ReceiverConfirmTime sql.NullTime `json:"receiver_confirm_time"`
	RequestTime         sql.NullTime `json:"request_time"`
}

type HeartbeatDetail struct {
	HeartbeatRequestId  uint         `json:"heartbeat_request_id"`
	SenderUserId        int64        `json:"sender_user_id"`
	SenderName          string       `json:"sender_name"`
	SenderAvatar        string       `json:"sender_avatar"`
	SenderGender        int8         `json:"sender_gender"`
	SenderIdentityTag   int8         `json:"sender_identity_tag"`
	ReceiverUserId      int64        `json:"receiver_user_id"`
	ReceiverName        string       `json:"receiver_name"`
	ReceiverAvatar      string       `json:"receiver_avatar"`
	ReceiverGender      int8         `json:"receiver_gender"`
	ReceiverIdentityTag int8         `json:"receiver_identity_tag"`
	MatchingStatus      int8         `json:"matching_status"`
	ReceiverConfirmTime sql.NullTime `json:"receiver_confirm_time"`
	RequestTime         sql.NullTime `json:"request_time"`
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
