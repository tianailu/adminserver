package models

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/plugin/soft_delete"
	"log"
	"time"
)

type (
	User struct {
		Id             uint                  `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		AccountId      string                `json:"account_id" gorm:"size:32;comment:账号ID;index:idx_account_id"`
		UserId         int64                 `json:"user_id" gorm:"not null;comment:用户ID;index:idx_uid"`
		Name           string                `json:"name" gorm:"size:8;comment:用户名;index:idx_name"`
		Avatar         string                `json:"avatar" gorm:"size:128;comment:个人头像"`
		Gender         int8                  `json:"gender" gorm:"not null;default:1;comment:性别，取值为[0:未选择, 1:男, 2:女]"`
		Birthday       sql.NullTime          `json:"birthday" gorm:"type:datetime;comment:出生日期"`
		Constellation  string                `json:"constellation" gorm:"size:12;comment:星座"`
		Height         float32               `json:"height" gorm:"default:0.0;comment:身高，单位cm"`
		Weight         float32               `json:"weight" gorm:"default:0.0;comment:体重，单位kg"`
		Education      int8                  `json:"education" gorm:"default:0;comment:最高学历，取值为[0:未选择, 1:博士及以上, 2:硕士, 3:本科, 4:专科, 5:高中及以下]"`
		EduStatus      int8                  `json:"edu_status" gorm:"default:0;comment:学历状态，取值为[0:未选择, 1:在校学生, 2:已毕业]"`
		School         string                `json:"school" gorm:"size:32;comment:毕业院校"`
		Work           int                   `json:"work" gorm:"default=0;comment:职业"`
		Company        string                `json:"company" gorm:"size:20;comment:公司"`
		Income         int8                  `json:"income" gorm:"default=0;comment:年收入，取值为[0:未选择, 1:5-10万, 2:11-20万, 3:21-30万, 4:31-50万, 5:51-100万, 6:101-200万, 7:201-500, 8:501-1000万, 9:1000万+]"`
		Residence      string                `json:"residence" gorm:"size:12;comment:现居住地（国家地理编码）"`
		Hometown       string                `json:"hometown" gorm:"size:12;comment:家乡（国家地理编码）"`
		MobilePhone    string                `json:"mobile_phone" gorm:"size:12;comment:手机号码"`
		IdentityTag    int8                  `json:"identity_tag" gorm:"not null;default=0;comment:身份标签，取值为[0:未选择, 1:母胎单身, 2:未婚单身, 3:离异无孩, 4:离异带孩, 5:离异不带孩, 6:恋爱中, 7:即将分手中]"`
		IsVip          int8                  `json:"is_vip" gorm:"not null;default:0;comment:是否vip，取值为[0:未知, 1:是, 2:否]"`
		VipTag         int8                  `json:"vip_tag" gorm:"not null;default:0;comment:vip标签"`
		Recommend      int8                  `json:"recommend" gorm:"not null;default:0;comment:推荐设置，取值为[0:未选择, 1:是, 2:否]"`
		RegisterPlace  string                `json:"register_place" gorm:"size:12;comment:注册地（国家地理编码）"`
		RegisterSource int8                  `json:"register_source" gorm:"comment:注册来源，取值为[0:未知, 1:APP, 2:小程序, 3:群组, 4:二维码, 5:管理后台]"`
		DurationOfUse  int64                 `json:"duration_of_use" gorm:"comment:使用时长，单位秒"`
		IsRealNameAuth int8                  `json:"is_rn_auth" gorm:"default:0;comment:是否完成实名认证，0:未认证，1:已通过认证"`
		IsWorkAuth     int8                  `json:"is_work_auth" gorm:"default:0;comment:是否完成工作认证，0:未认证，1:已通过认证"`
		IsEduAuth      int8                  `json:"is_edu_auth" gorm:"default:0;comment:是否完成学历认证，0:未认证，1:已通过认证"`
		AuditStatus    int8                  `json:"audit_status" gorm:"not null;default:1;comment:基础信息审核状态，取值为[0:未知, 1:待审（首次申请审核）, 2: 再审核（非首次申请审核）, 3:通过, 4:不通过]"`
		UserStatus     int8                  `json:"user_status" gorm:"not null;default:0;comment:用户状态，取值为[0:正常状态, 1:封号状态, 2:禁言状态, 3:注销状态]"`
		CreatedAt      time.Time             `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:注册时间"`
		UpdatedAt      time.Time             `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
		DeletedAt      sql.NullTime          `json:"deleted_at" gorm:"type:datetime;comment:注销时间"`
		IsDel          soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag;comment:删除标志，取值为[0:使用中, 1:已注销]"`
	}

	AboutMe struct {
		Id               uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId           int64     `json:"user_id" gorm:"not null;comment:用户唯一id;index:idx_user_id"`
		Habit            string    `json:"habit" gorm:"size:64;comment:生活习惯"`
		ConsumptionView  string    `json:"consumption_view" gorm:"size:64;comment:消费观"`
		FamilyBackground string    `json:"family_background" gorm:"size:64;comment:家庭背景"`
		Interest         string    `json:"interest" gorm:"size:64;comment:兴趣爱好"`
		LoveView         string    `json:"love_view" gorm:"size:64;comment:爱情观"`
		TargetAppearance string    `json:"ta_appearance" gorm:"size:64;comment:希望另一半的样子"`
		BeImpressed      string    `json:"be_impressed" gorm:"size:64;comment:对方什么最能打动自己"`
		CreatedAt        time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt        time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	MatchSetting struct {
		Id              uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId          int64     `json:"user_id" gorm:"not null;comment:用户唯一id;index:idx_user_id"`
		TargetAge       string    `json:"ta_age" gorm:"size:12;comment:希望另一半身高范围，中间使用英文横杠隔开，示例：18-38"`
		TargetHeight    string    `json:"ta_height" gorm:"size:12;comment:希望另一半身高范围，中间使用英文横杠隔开，示例170-190"`
		TargetCity      int8      `json:"ta_city" gorm:"default:0;comment:希望另一半所在城市，取值为[0:同城优先, 1:只要同城]"`
		TargetHometown  int8      `json:"ta_hometown" gorm:"default:0;comment:希望另一半的家乡，取值为[0:都可以, 1:同城优先]"`
		TargetEducation int8      `json:"ta_education" gorm:"default:0;comment:希望另一半最低学历，取值为[0:都可以, 1:本科, 2:硕士]"`
		TargetMarriage  int8      `json:"ta_marriage" gorm:"default:0;comment:希望另一半婚姻状态，取值为[0:未婚, 1:可以离异]"`
		CreatedAt       time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt       time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	RealNameAuth struct {
		Id        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId    int64     `json:"user_id" gorm:"not null;comment:用户唯一id;index:idx_user_id"`
		IdCard    string    `json:"id_card" gorm:"size:18;not null;comment:身份证"`
		RealName  string    `json:"real_name" gorm:"size:20;not null;comment:真实姓名"`
		Status    int8      `json:"status" gorm:"not null;default:0;comment:认证状态，取值为[0:未认证, 1:已通过认证, 2:认证未通过, 3:更新认证]"`
		CreatedAt time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	WorkAuth struct {
		Id                uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId            int64     `json:"user_id" gorm:"not null;comment:用户唯一id;index:idx_user_id"`
		AuthMethod        int8      `json:"auth_method" gorm:"comment:认证方式，取值为[0:未选择, 1:支付宝社保截图／社保证明, 2:钉钉或企业微名片（需带二维码）, 3:在职证明／劳动合同／营业执照, 4:工牌／名片／工作证等, 5:录取Offer／工资单]"`
		Company           string    `json:"company" gorm:"size:20;not null;comment:公司名称"`
		Img               string    `json:"img" gorm:"size:128;not null;comment:提供的认证图片"`
		IsBlockColleagues int8      `json:"is_block_colleagues" gorm:"default=0;comment:是否屏蔽同事，取值为[0:不屏蔽, 1:屏蔽]"`
		Status            int8      `json:"status" gorm:"not null;default:0;comment:认证状态，取值为[0:未认证, 1:已通过认证, 2:认证未通过, 3:更新认证]"`
		CreatedAt         time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt         time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	EduAuth struct {
		Id               uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId           int64     `json:"user_id" gorm:"not null;comment:用户唯一id;index:idx_user_id"`
		SchoolType       int8      `json:"school_type" gorm:"default:0;comment:学校类型，取值为[0:内地学校, 1:海外/港澳台]"`
		AuthMethod       int8      `json:"auth_method" gorm:"comment:认证方式，取值为[0:未选择, 1:内地在线自助认证, 2:内地毕业证书/学位证书编码, 3:内地毕业证书/学位证书照片, 4:学信网在线验证码, 5:教留服认证证书编号, 6:海外/港澳台学历证书照片, 7:在校学生证明]"`
		Education        int8      `json:"education" gorm:"default:0;comment:最高学历，取值为[0:未选择, 1:高中及以下, 2:大专, 3:本科, 4:硕士, 5:博士及以上]"`
		School           string    `json:"school" gorm:"size:32;comment:学校名称"`
		CertificateNo    string    `json:"cert_no" gorm:"size:20;comment:证书号码"`
		VerificationCode string    `json:"verification_code" gorm:"size:20;comment:学信网在线验证码"`
		Img              string    `json:"img" gorm:"size:128;not null;comment:提供的认证图片"`
		Status           int8      `json:"status" gorm:"not null;default:0;comment:认证状态，取值为[0:未认证, 1:已通过认证, 2:认证未通过, 3:更新认证]"`
		CreatedAt        time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt        time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	UserManagement struct {
		Id        uint         `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId    int64        `json:"user_id" gorm:"not null;comment:用户唯一id;index:idx_user_id"`
		Action    int8         `json:"action" gorm:"not null;comment:操作类型，取值为[1:封号, 2:禁言]"`
		Reason    int8         `json:"reason" gorm:"default:0;comment:操作理由，取值为[0:未选择, 1:垃圾营销广告, 2:色情低俗, 3:政治敏感, 4:虚假信息, 5:资料透露联系方式, 6:聊天内容不适]"`
		Penalties int8         `json:"penalties" gorm:"default:0;comment:处罚措施，取值为[0:无限期, 1:1天, 2:2天, 3:3天, 4:7天, 5:10天]"`
		Until     sql.NullTime `json:"until" gorm:"type:datetime;comment:解封时间"`
		Remark    string       `json:"remark" gorm:"comment:备注"`
		CreatedAt time.Time    `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt time.Time    `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	Reports struct {
		Id             uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		ReporterUserId int64     `json:"reporter_user_id" gorm:"not null;comment:举报人用户ID;index:idx_reporter_user_id"`
		ReportedUserId int64     `json:"reported_user_id" gorm:"not null;comment:被举报人用户ID;index:idx_reported_user_id"`
		ReportSource   int8      `json:"report_source" gorm:"comment:举报来源，取值为[]"`
		ReportType     int8      `json:"report_type" gorm:"default:0;comment:举报事项类型，取值为[0:其他, 1:头像非本人, 2:资料透露练习方式, 3:内容乱填/虚假资料, 4:婚托/酒托/饭托等, 5:虚假中奖消息、诈骗等, 6:垃圾营销广告, 7:聊天内容不适/骚扰]"`
		Desc           string    `json:"desc" gorm:"type:text;comment:具体描述"`
		Img            string    `json:"img" gorm:"举报图片，多张图用英文逗号隔开"`
		Status         int8      `json:"status" gorm:"default:0;comment:状态，取值为[0:未处理, 1:已处理]"`
		ReportTime     time.Time `json:"report_time" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:举报时间"`
		UpdatedAt      time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	VipTag struct {
		Id                    uint            `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		Name                  string          `json:"name" gorm:"size:10;not null;comment:标签名称"`
		GrossTransactionValue decimal.Decimal `json:"gross_transaction_value" gorm:"type:decimal(10,2);comment:累计交易金额 GTV"`
		CreatedAt             time.Time       `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt             time.Time       `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	UserVipTag struct {
		Id        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId    int64     `json:"user_id" gorm:"not null;comment:用户唯一id;index:idx_user_id"`
		VipTagId  uint      `json:"vip_tag_id" gorm:"not null;comment:会员标签id"`
		CreatedAt time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	Follow struct {
		Id             uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId         int64     `json:"user_id" gorm:"not null;comment:用户id;index:idx_user_id"`
		FollowedUserId int64     `json:"follow_user_id" gorm:"not null;comment:被关注人的用户id"`
		CreatedAt      time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	}

	Fans struct {
		Id         uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId     int64     `json:"user_id" gorm:"not null;comment:用户id;index:idx_user_id"`
		FansUserId int64     `json:"fans_user_id" gorm:"not null;comment:粉丝的用户id"`
		CreatedAt  time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	}

	Friendship struct {
		Id        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		User1Id   int64     `json:"user_1_id" gorm:"column:user_1_id;not null;comment:用户1的id，保证 user_1_id < user_2_id;index:idx_user_1_id"`
		User2Id   int64     `json:"user_2_id" gorm:"column:user_2_id;not null;comment:用户2的id，保证 user_1_id < user_2_id;index:idx_user_2_id"`
		Status    int8      `json:"status" gorm:"not null;comment:好友关系状态，取值为[1:user1发出的申请, 2:user2发出的申请, 3:user1拉黑对方, 4:user2拉黑对方, 5:互相拉黑]"`
		CreatedAt time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	FriendRequest struct {
		Id                  uint         `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		SenderUserId        int64        `json:"sender_user_id" gorm:"not null;comment:发起好友申请的用户id;index:idx_sender_id"`
		ReceiverUserId      int64        `json:"receiver_user_id" gorm:"not null;comment:接收好友申请的用户id;index:idx_receiver_id"`
		Question            string       `json:"question" gorm:"size:64;comment:灵魂交友问题"`
		Answer              string       `json:"answer" gorm:"size:64;comment:回答"`
		MatchingStatus      int8         `json:"matching_status" gorm:"not null;default=1;comment:好友申请状态，取值为[1:待确认处理, 2:已接受, 3:被拒绝, 4:主动中止申请, 5:再次申请]"`
		ReceiverReadTime    sql.NullTime `json:"receiver_read_time" gorm:"type:datetime;comment:被申请人首次查看到申请的时间"`
		ReceiverConfirmTime sql.NullTime `json:"receiver_confirm_time" gorm:"type:datetime;comment:被申请人回复申请时间"`
		CreatedAt           time.Time    `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt           time.Time    `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	HeartbeatMatching struct {
		Id        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		User1Id   int64     `json:"user_1_id" gorm:"column:user_1_id;not null;comment:用户1的id，保证 user_1_id < user_2_id;index:idx_user_1_id"`
		User2Id   int64     `json:"user_2_id" gorm:"column:user_2_id;not null;comment:用户2的id，保证 user_1_id < user_2_id;index:idx_user_2_id"`
		Status    int8      `json:"status" gorm:"not null;comment:好友关系状态，取值为[1:user1发出的申请, 2:user2发出的申请]"`
		CreatedAt time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	HeartbeatRequest struct {
		Id                  uint         `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		SenderUserId        int64        `json:"sender_user_id" gorm:"not null;comment:发起心动匹配申请的用户id;index:idx_sender_id"`
		ReceiverUserId      int64        `json:"receiver_user_id" gorm:"not null;comment:接收到心动匹配申请的用户id;index:idx_receiver_id"`
		MatchingStatus      int8         `json:"matching_status" gorm:"not null;default=1;comment:心动匹配申请状态，取值为[1:待确认处理, 2:已接受]"`
		ReceiverReadTime    sql.NullTime `json:"receiver_read_time" gorm:"type:datetime;comment:被申请人首次查看到申请的时间"`
		ReceiverConfirmTime sql.NullTime `json:"receiver_confirm_time" gorm:"type:datetime;comment:被申请人回复申请时间"`
		CreatedAt           time.Time    `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt           time.Time    `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	FindCompanionActivity struct {
		Id               uint         `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId           int64        `json:"user_id" gorm:"not null;comment:用户id;index:idx_user_id"`
		CompanionTypeId  int          `json:"companion_type_id" gorm:"not null;comment:搭子类型id"`
		ActivityName     string       `json:"activity_name" gorm:"size:800;comment:活动名称"`
		ActivityTime     time.Time    `json:"activity_time" gorm:"type:datetime;not null;comment:活动时间"`
		ActivityLocation string       `json:"activity_location" gorm:"comment:活动地点"`
		Latitude         float32      `json:"latitude" gorm:"comment:活动地点-维度"`
		Longitude        float32      `json:"longitude" gorm:"comment:活动地点-经度"`
		CostType         int8         `json:"cost_type" gorm:"comment:活动费用类型，取值为[0:未选择, 1:发起人请客, 2:AA, 3:对方请客]"`
		Desc             string       `json:"desc" gorm:"size:800;comment:描述/想说"`
		Status           int8         `json:"status" gorm:"comment:活动状态，取值为[1:进行中, 2:匹配成功, 3:已取消, 4:已结束]"`
		CancelTime       sql.NullTime `json:"cancel_time" gorm:"type:datetime;comment:取消活动的时间"`
		CreatedAt        time.Time    `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt        time.Time    `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	FindCompanionRequest struct {
		Id              uint         `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		ActivityId      uint         `json:"activity_id" gorm:"not null;comment:活动ID;index:idx_activity_id"`
		SenderUserId    int64        `json:"sender_user_id" gorm:"not null;comment:发起找搭子活动的用户id;index:idx_sender_id"`
		ApplicantUserId int64        `json:"applicant_user_id" gorm:"not null;comment:申请参与活动的用户id;index:idx_applicant_id"`
		ApplicationTime time.Time    `json:"application_time" gorm:"type:datetime;not null;comment:申请参与活动的时间"`
		Status          int8         `json:"status" gorm:"not null;comment:申请状态，取值为[1:申请待处理中, 2:已通过申请, 3:已取消申请]"`
		CancelTime      sql.NullTime `json:"cancel_time" gorm:"type:datetime;comment:申请人取消参与活动的时间"`
		CreatedAt       time.Time    `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt       time.Time    `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	CompanionType struct {
		Id        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		Tag       int8      `json:"tag" gorm:"not null;comment:主标签，取值为[1:美食搭子, 2:日常娱乐搭子, 3:户外/旅行搭子, 4:运动/健身搭子, 5:学习/进步搭子]"`
		Name      string    `json:"name" gorm:"size:12;not null;comment:搭子类型名称"`
		Status    int8      `json:"status" gorm:"not null;comment:状态，取值为[1:被选择，2:未选择]"`
		CreatedAt time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}
)

func (m *User) TableName() string {
	return "tb_user"
}

func (m *AboutMe) TableName() string {
	return "tb_about_me"
}

func (m *MatchSetting) TableName() string {
	return "tb_match_setting"
}

func (m *RealNameAuth) TableName() string {
	return "tb_real_name_auth"
}

func (m *WorkAuth) TableName() string {
	return "tb_work_auth"
}

func (m *EduAuth) TableName() string {
	return "tb_edu_auth"
}

func (m *UserManagement) TableName() string {
	return "tb_user_management"
}

func (m *Reports) TableName() string {
	return "tb_reports"
}

func (m *VipTag) TableName() string {
	return "tb_vip_tag"
}

func (m *UserVipTag) TableName() string {
	return "tb_user_vip_tag"
}

func (m *Follow) TableName() string {
	return "tb_follow"
}

func (m *Fans) TableName() string {
	return "tb_fans"
}

func (m *Friendship) TableName() string {
	return "tb_friendship"
}

func (m *FriendRequest) TableName() string {
	return "tb_friend_request"
}

func (m *HeartbeatMatching) TableName() string {
	return "tb_heartbeat_matching"
}

func (m *HeartbeatRequest) TableName() string {
	return "tb_heartbeat_request"
}

func (m *FindCompanionActivity) TableName() string {
	return "tb_find_companion_activity"
}

func (m *FindCompanionRequest) TableName() string {
	return "tb_find_companion_request"
}

func (m *CompanionType) TableName() string {
	return "tb_companion_type"
}

func CreateTable() error {
	err := mysql.GetDB().Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(&User{}, &AboutMe{}, &MatchSetting{}, &RealNameAuth{}, &WorkAuth{}, &EduAuth{}, &UserManagement{},
			&Reports{}, &VipTag{}, &UserVipTag{}, &Follow{}, &Fans{}, &Friendship{}, &FriendRequest{},
			&HeartbeatMatching{}, &HeartbeatRequest{})
	if err != nil {
		log.Printf("创建 tb_user/tb_about_me/tb_match_setting/tb_real_name_auth/tb_work_auth/tb_edu_auth/tb_user_management/tb_reports/tb_vip_tag/tb_user_vip_tag 表失败, err: %s", err)
		return err
	}

	// 设置表备注
	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(User).TableName(), "用户详情表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(User).TableName(), err)
		return err
	}

	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(AboutMe).TableName(), "用户个人信息介绍表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(AboutMe).TableName(), err)
		return err
	}

	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(MatchSetting).TableName(), "匹配设置表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(MatchSetting).TableName(), err)
		return err
	}

	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(RealNameAuth).TableName(), "实名认证表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(RealNameAuth).TableName(), err)
		return err
	}

	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(WorkAuth).TableName(), "工作认证表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(WorkAuth).TableName(), err)
		return err
	}

	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(EduAuth).TableName(), "学历认证表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(EduAuth).TableName(), err)
		return err
	}

	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(UserManagement).TableName(), "用户管理表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(UserManagement).TableName(), err)
		return err
	}
	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(Reports).TableName(), "举报管理表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(Reports).TableName(), err)
		return err
	}
	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(VipTag).TableName(), "会员标签表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(VipTag).TableName(), err)
		return err
	}
	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(UserVipTag).TableName(), "用户会员信息表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(UserVipTag).TableName(), err)
		return err
	}
	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(Follow).TableName(), "用户关注记录表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(Follow).TableName(), err)
		return err
	}
	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(Fans).TableName(), "用户粉丝记录表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(Fans).TableName(), err)
		return err
	}
	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(Friendship).TableName(), "用户好友关系表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(Friendship).TableName(), err)
		return err
	}
	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(FriendRequest).TableName(), "用户好友申请记录表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(FriendRequest).TableName(), err)
		return err
	}
	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(HeartbeatMatching).TableName(), "心动匹配关系表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(HeartbeatMatching).TableName(), err)
		return err
	}
	err = mysql.GetDB().Exec(fmt.Sprintf("ALTER TABLE %s COMMENT = '%s'", new(HeartbeatRequest).TableName(), "心动匹配请求记录表")).Error
	if err != nil {
		log.Printf("添加表备注失败, table: %s, err: %s", new(HeartbeatRequest).TableName(), err)
		return err
	}

	return nil
}
