package models

type UserSearchParam struct {
	Keywords        string `query:"keywords"`                               // 关键字，用户ID/昵称/用户名
	Gender          int8   `query:"gender" validate:"lte=2,gte=0"`          // 性别，取值为[0:全部, 1:男, 2:女]，默认值为0。
	IdentityTag     int8   `query:"identity_tag" validate:"lte=5,gte=0"`    // 身份标签，取值为[0:全部, 1:母胎单身, 2:未婚单身, 3:离异无孩, 4:离异带孩, 5:离异不带孩]，默认值为0。
	IsVip           int8   `query:"is_vip" validate:"lte=2,gte=0"`          // 是否vip，取值为[0:全部, 1:是, 2:否]，默认值为0。
	VipTag          int32  `query:"vip_tag" validate:"gte=0"`               // VIP标签Id
	AuditStatus     int8   `query:"audit_status" validate:"lte=4,gte=0"`    // 基础信息审核状态，取值为[0:全部, 1:待审（首次申请审核）, 2: 再审核（非首次申请审核）, 3:通过, 4:不通过]，默认值为0。
	Income          int8   `query:"income" validate:"lte=9,gte=0"`          // 年收入，取值为[0:全部, 1:5-10万, 2:11-20万, 3:21-30万, 4:31-50万, 5:51-100万, 6:101-200万, 7:201-500, 8:501-1000万, 9:1000万+]，默认值为0。
	Recommend       int8   `query:"recommend" validate:"lte=2,gte=0"`       // 推荐，取值为[0:全部, 1:是, 2:否]，默认值为0。
	RegisterPlace   string `query:"register_place"`                         // 注册地
	RegisterSource  int8   `query:"register_source" validate:"lte=5,gte=0"` // 注册来源，取值为[0:全部, 1:APP, 2:小程序, 3:群组, 4:二维码, 5:管理后台]，默认值为0。
	RegisterStartAt int64  `query:"register_start_at"`                      // 开始时间，时间戳，单位毫秒
	RegisterEndAt   int64  `query:"register_end_at"`                        // 结束时间，时间戳，单位毫秒
	PageNum         int    `query:"page_num" validate:"gte=0"`              // 页码，默认值为1。
	PageSize        int    `query:"page_size" validate:"gte=0"`             // 每页大小，默认值为20。
}

type UserListItem struct {
	UserId                         int64   `json:"user_id,optional"`                            // 用户ID
	Name                           string  `json:"name,optional"`                               // 昵称
	Gender                         int8    `json:"gender,optional"`                             // 性别
	AuditStatus                    int8    `json:"audit_status,optional"`                       // 审核状态
	IdentityTag                    int8    `json:"identity_tag,optional"`                       // 身份标签/情感状态
	IsVip                          int8    `json:"is_vip,optional"`                             // 是否Vip
	VipTag                         int32   `json:"vip_tag,optional"`                            // Vip标签
	RechargeAmount                 float32 `json:"recharge_amount,optional"`                    // 充值金额
	RemainingCoins                 int64   `json:"remaining_coins,optional"`                    // 剩余金币
	ConsumeCoins                   int64   `json:"consume_coins,optional"`                      // 消耗金币
	Income                         int8    `json:"income,optional"`                             // 年收入
	RegisterPlace                  string  `json:"register_place,optional"`                     // 注册地
	RegisterSource                 int8    `json:"register_source,optional"`                    // 注册来源
	RegisterTime                   int64   `json:"register_time,optional"`                      // 注册时间
	TotalUsageTime                 int64   `json:"total_usage_time,optional"`                   // 使用时长
	FriendRequestCount             int64   `json:"friend_request_count,optional"`               // 申请认识数量/发起好友申请数量
	FriendRequestSuccessCount      int64   `json:"friend_request_success_count,optional"`       // 申请认识成功数量/成为好友数量
	HeartbeatRequestCount          int64   `json:"heartbeat_request_count,optional"`            // 发起心动匹配数量
	HeartbeatRequestSuccessCount   int64   `json:"heartbeat_request_success_count,optional"`    // 心动匹配成功次数
	FindCompanionRequestCount      int64   `json:"find_companion_request_count,optional"`       // 找搭子次数
	FindCompanionSuccessMatchCount int64   `json:"find_companion_success_match_count,optional"` // 找搭子匹配成功次数
	FollowCount                    int64   `json:"follow_count,optional"`                       // 关注用户数量
	FansCount                      int64   `json:"fans_count,optional"`                         // 被关注数量/粉丝数量
	ShareCount                     int64   `json:"share_count,optional"`                        // 发起分享次数
	// TODO 交友相关数据
}

type UserSimple struct {
	UserId int64  `json:"user_id,optional"`
	Name   string `json:"name,optional"`
	Gender int8   `json:"gender,optional"`
}

type LogoutUser struct {
	UserId   int64  `json:"user_id,optional"`
	Name     string `json:"name,optional"`
	Gender   int8   `json:"gender,optional"`
	DeleteAt int64  `json:"delete_at"`
}

type FriendSearchParam struct {
	Keywords           string `query:"keywords,optional"`             // 关键字，用户ID/昵称/用户名
	MatchType          int8   `query:"match_type,optional"`           // 匹配类型，取值为[0:全部, 1:发出的申请, 2:接收到的申请]，默认值为0。
	MatchingStatus     int8   `query:"matching_status,optional"`      // 匹配状态，取值为[0:全部, 1:待确认, 2:已接受, 3:被拒绝, 4:主动中止申请]，默认值为0。
	ApplicationStartAt int64  `query:"application_start_at,optional"` // 开始申请时间，时间戳，单位毫秒
	ApplicationEndAt   int64  `query:"application_end_at,optional"`   // 结束申请时间，时间戳，单位毫秒
	PageNum            int    `query:"page_num,optional"`             // 页码，默认值为1。
	PageSize           int    `query:"page_size,optional"`            // 每页大小，默认值为20。
}

type FriendListItem struct {
	FriendRequestId     int64  `json:"friend_request_id"`
	SenderUserId        int64  `json:"sender_user_id"`
	SenderName          string `json:"sender_name"`
	SenderAvatar        string `json:"sender_avatar"`
	SenderGender        int8   `json:"sender_gender"`
	SenderIdentityTag   int8   `json:"sender_identity_tag"`
	RequestTime         int64  `json:"request_time"`
	ReceiverUserId      int64  `json:"receiver_user_id"`
	ReceiverName        string `json:"receiver_name"`
	ReceiverAvatar      string `json:"receiver_avatar"`
	ReceiverGender      int8   `json:"receiver_gender"`
	ReceiverIdentityTag int8   `json:"receiver_identity_tag"`
	ReceiverConfirmTime int64  `json:"receiver_confirm_time"`
	MatchingStatus      int8   `json:"matching_status"`
}

type HeartbeatSearchParam struct {
	Keywords           string `query:"keywords,optional"`             // 关键字，用户ID/昵称/用户名
	MatchType          int8   `query:"match_type,optional"`           // 匹配类型，取值为[0:全部, 1:发出的申请, 2:接收到的申请]，默认值为0。
	MatchingStatus     int8   `query:"matching_status,optional"`      // 匹配状态，取值为[0:全/部, 1:待确认, 2:已接受]，默认值为0。
	ApplicationStartAt int64  `query:"application_start_at,optional"` // 开始申请时间
	ApplicationEndAt   int64  `query:"application_end_at,optional"`   // 结束申请时间
	PageNum            int    `query:"page_num,optional"`             // 页码，默认值为1。
	PageSize           int    `query:"page_size,optional"`            // 每页大小，默认值为20。
}

type HeartbeatListItem struct {
	HeartbeatRequestId  int64  `json:"heartbeat_request_id"`
	SenderUserId        int64  `json:"sender_user_id"`
	SenderName          string `json:"sender_name"`
	SenderAvatar        string `json:"sender_avatar"`
	SenderGender        int8   `json:"sender_gender"`
	SenderIdentityTag   int8   `json:"sender_identity_tag"`
	RequestTime         int64  `json:"request_time"`
	ReceiverUserId      int64  `json:"receiver_user_id"`
	ReceiverName        string `json:"receiver_name"`
	ReceiverAvatar      string `json:"receiver_avatar"`
	ReceiverGender      int8   `json:"receiver_gender"`
	ReceiverIdentityTag int8   `json:"receiver_identity_tag"`
	ReceiverConfirmTime int64  `json:"receiver_confirm_time"`
	MatchingStatus      int8   `json:"matching_status"`
}

type FindCompanionSearchParam struct {
	Keywords       string `query:"keywords,optional"`         // 关键字，用户ID/昵称/用户名
	CompanionTag   int8   `query:"companion_tag"`             // 搭子主标签，取值为[1:美食搭子, 2:日常娱乐搭子, 3:户外/旅行搭子, 4:运动/健身搭子, 5:学习/进步搭子]
	MatchType      int8   `query:"match_type,optional"`       // 匹配类型，取值为[0:全部, 1:发起搭子活动, 2:报名搭子活动]，默认值为0。
	MatchingStatus int8   `query:"matching_status,optional"`  // 匹配状态，取值为[0:全部, 1:进行中, 2:匹配成功, 3:已取消]，默认值为0。
	PublishStartAt int64  `query:"publish_start_at,optional"` // 活动发布时间-开始
	PublishEndAt   int64  `query:"publish_end_at,optional"`   // 活动发布时间-结束
	PageNum        int    `query:"page_num,optional"`         // 页码，默认值为1。
	PageSize       int    `query:"page_size,optional"`        // 每页大小，默认值为20。
}

type FindCompanionListItem struct {
	FindCompanionActivityId int64  `json:"find_companion_activity_id"`
	SenderUserId            int64  `json:"sender_user_id"`
	SenderName              string `json:"sender_name"`
	SenderAvatar            string `json:"sender_avatar"`
	SenderGender            int8   `json:"sender_gender"`
	SenderIdentityTag       int8   `json:"sender_identity_tag"`
	ReceiverUserId          int64  `json:"receiver_user_id"`
	ReceiverName            string `json:"receiver_name"`
	ReceiverAvatar          string `json:"receiver_avatar"`
	ReceiverGender          int8   `json:"receiver_gender"`
	ReceiverIdentityTag     int8   `json:"receiver_identity_tag"`
	CompanionTypeId         int    `json:"companion_type_id"`
	CompanionTypeName       string `json:"companion_type_name"`
	MatchingStatus          int8   `json:"matching_status"`
	PublishingActivityTime  int64  `json:"publishing_activity_time"`
	RequestTime             int64  `json:"request_time"`
	ActivityCancelTime      int64  `json:"activity_cancel_time"`
	SenderConfirmTime       int64  `json:"sender_confirm_time"`
}

type CompanionTypeSearchParam struct {
	Status int8 `query:"status,optional"` // 状态，取值为[0:全部, 1:被选择，2:未选择]
}

type CompanionTypeListItem struct {
	CompanionTypeId   int64  `json:"companion_type_id"`
	CompanionTag      int8   `json:"companion_tag"`
	CompanionTypeName string `json:"companion_type_name"`
	Status            int8   `json:"status"`
	CreateAt          int64  `json:"create_at"`
}
