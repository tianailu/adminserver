package user

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"log"
	"time"
)

type (
	User struct {
		Id             uint         `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		AccountId      string       `json:"account_id" gorm:"size:32;comment:账号ID"`
		Uid            string       `json:"uid" gorm:"size:32;not null;comment:用户ID"`
		Name           string       `json:"name" gorm:"size:12;comment:用户名"`
		Avatar         string       `json:"avatar" gorm:"size:128;comment:个人头像"`
		Gender         int8         `json:"gender" gorm:"not null;default:0;comment:性别，取值为[0:男, 1:女]"`
		Birthday       sql.NullTime `json:"birthday" gorm:"type:datetime;comment:出生日期"`
		Constellation  string       `json:"constellation" gorm:"size:12;comment:星座"`
		Height         float32      `json:"height" gorm:"default:0.0;comment:身高，单位cm"`
		Weight         float32      `json:"weight" gorm:"default:0.0;comment:体重，单位kg"`
		Education      string       `json:"education" gorm:"size:12;comment:最高学历"`
		School         string       `json:"school" gorm:"size:32;comment:毕业院校"`
		Work           string       `json:"work" gorm:"size:12;comment:职业"`
		Company        string       `json:"company" gorm:"size:20;comment:公司"`
		Income         string       `json:"income" gorm:"size:20;comment:年收入"`
		Residence      string       `json:"residence" gorm:"size:12;comment:现居住地（国家地理编码）"`
		Hometown       string       `json:"hometown" gorm:"size:12;comment:家乡（国家地理编码）"`
		MobilePhone    string       `json:"mobile_phone" gorm:"size:12;comment:手机号码"`
		IdentityTag    int8         `json:"identity_tag" gorm:"not null;default=0;comment:身份标签，取值为[0:未选择, 1:母胎单身, 2:未婚单身, 3:离异无孩, 4:离异带孩, 5:离异不带孩, 6:丧偶]"`
		VipTag         int8         `json:"vip_tag" gorm:"not null;default:0;comment:vip标签"`
		Status         int8         `json:"status" gorm:"not null;default:0;comment:用户状态，取值为[0:待审（首次申请审核）, 1: 再审核（非首次申请审核）, 2:通过, 3:不通过]"`
		RegisterPlace  string       `json:"register_place" gorm:"size:12;comment:注册地（国家地理编码）"`
		RegisterSource int8         `json:"register_source" gorm:"comment:注册来源，取值为[0:未知, 1:APP, 2:小程序, 3:群组, 4:二维码, 5:管理后台]"`
		DurationOfUse  int64        `json:"duration_of_use" gorm:"comment:使用时长，单位秒"`
		CreatedAt      time.Time    `json:"created_at" gorm:"comment:注册时间"`
		UpdateAt       time.Time    `json:"update_at" gorm:"comment:修改时间"`

		WeChat   string `json:"wechat" gorm:"size:20;comment:微信号"`
		Marriage string `json:"marriage" gorm:"size:12;comment:婚姻状态"`

		IdCard         string `json:"id_card" gorm:"size:18;comment:身份证"`
		RealName       string `json:"real_name" gorm:"size:20;comment:真实姓名"`
		IsRealNameAuth int8   `json:"is_rn_auth" gorm:"default:0;comment:是否完成实名认证，0:未认证，1:已通过认证"`
		WorkAuthImg    string `json:"work_auth_img" gorm:"size:128;comment:工作认证图片"`
		IsWorkAuth     int8   `json:"is_work_auth" gorm:"default:0;comment:是否完成工作认证，0:未认证，1:已通过认证"`
		EduAuthImg     string `json:"edu_auth_img" gorm:"size:128;comment:学历认证图片"`
		IsEduAuth      int8   `json:"is_edu_auth" gorm:"default:0;comment:是否完成学历认证，0:未认证，1:已通过认证"`
	}

	AboutMe struct {
		Id               uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId           uint      `json:"user_id" gorm:"comment:用户唯一id"`
		Habit            string    `json:"habit" gorm:"size:64;comment:生活习惯"`
		ConsumptionView  string    `json:"consumption_view" gorm:"size:64;comment:消费观"`
		Family           string    `json:"family" gorm:"size:64;comment:家庭背景"`
		Interest         string    `json:"interest" gorm:"size:64;comment:兴趣爱好"`
		LoveView         string    `json:"love_view" gorm:"size:64;comment:爱情观"`
		TargetAppearance string    `json:"ta_appearance" gorm:"size:64;comment:希望另一半的样子"`
		BeImpressed      string    `json:"be_impressed" gorm:"size:64;comment:对方什么最能打动自己"`
		CreatedAt        time.Time `json:"created_at" gorm:"comment:创建时间"`
		UpdateAt         time.Time `json:"update_at" gorm:"comment:修改时间"`
	}

	MatchSetting struct {
		Id              uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId          uint      `json:"user_id" gorm:"comment:用户唯一id"`
		TargetAge       string    `json:"ta_age" gorm:"size:12;comment:希望另一半身高范围，中间使用英文横杠隔开，示例：18-38"`
		TargetHeight    string    `json:"ta_height" gorm:"size:12;comment:希望另一半身高范围，中间使用英文横杠隔开，示例170-190"`
		TargetCity      int8      `json:"ta_city" gorm:"default:0;comment:希望另一半所在城市，取值为[0:同城优先, 1:只要同城]"`
		TargetHometown  int8      `json:"ta_hometown" gorm:"default:0;comment:希望另一半的家乡，取值为[0:都可以, 1:同城优先]"`
		TargetEducation int8      `json:"ta_education" gorm:"default:0;comment:希望另一半最低学历，取值为[0:都可以, 1:本科, 2:硕士]"`
		TargetMarriage  int8      `json:"ta_marriage" gorm:"default:0;comment:希望另一半婚姻状态，取值为[0:未婚, 1:可以离异]"`
		CreatedAt       time.Time `json:"created_at" gorm:"comment:创建时间"`
		UpdateAt        time.Time `json:"update_at" gorm:"comment:修改时间"`
	}

	VipLevel struct {
		Id                  uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		Level               string    `json:"level" gorm:"size:10;not null;comment:等级"`
		Name                string    `json:"name" gorm:"size:10;not null;comment:等级名称"`
		Desc                string    `json:"desc" gorm:"size:90;comment:描述"`
		Weights             int8      `json:"weights" gorm:"default=0;comment:等级权重，取值为[0:未选择, 1:完全公开, 2:私密, 3:仅好友]"`
		ProductId           uint      `json:"product_id" gorm:"comment:商品id"`
		DiscountStatus      int8      `json:"discount_status" gorm:"default=0;comment:等级权益（折扣）开关，取值为[0:关闭, 1:开启]"`
		Discount            float32   `json:"discount" gorm:"default=1;comment:等级权益（折扣）"`
		UpgradeStrategyId   uint      `json:"upgrade_strategy_id" gorm:"comment:升级策略Id"`
		DowngradeStrategyId uint      `json:"downgrade_strategy_id" gorm:"comment:降级策略Id"`
		Status              int8      `json:"status" gorm:"comment:状态，取值为[0:禁用, 1:启用]"`
		CreatedAt           time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt           time.Time `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	VipStrategy struct {
		Id               uint              `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		Type             int8              `json:"type" gorm:"not null;comment:策略类型，取值为[1:升级, 2:降级]"`
		ProtectionPeriod int               `json:"protection_period" gorm:"default=0;not null;comment:等级保护期，单位为天，值为0时无视"`
		ConditionType    int8              `json:"condition_type" gorm:"not null;comment:策略达成条件类型，取值为[1: 满足任意一个条件, 2:满足全部条件]"`
		AssessmentPeriod int               `json:"assessment_period" gorm:"not null;comment:考核周期，考核最近一个周期内的数据，单为位天，值为0时考核过去所有的累计数据"`
		Condition        StrategyCondition `json:"condition" gorm:"type:json;default={};comment:考核条件，json字符串"`
		Status           int8              `json:"status" gorm:"comment:策略开关，取值为[0:关闭, 1:开启]"`
		StartTime        sql.NullTime      `json:"start_time" gorm:"type:datetime;comment:策略开始启用时间，为空时即时生效"`
		CreatedAt        time.Time         `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt        time.Time         `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	StrategyCondition map[string]interface{}

	VipTag struct {
		Id                    uint            `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		Name                  string          `json:"name" gorm:"size:10;comment:标签名称"`
		AutoTagStatus         int8            `json:"auto_tag_status" gorm:"default=0;comment:自动打标签状态，取值为[0:关闭, 1:启用]"`
		StrategyCondition     int8            `json:"strategy_condition" gorm:"default=0;comment:打标签策略类型，取值为[0:满足任意一个条件, 1:满足全部条件]"`
		GrossTransactionValue decimal.Decimal `json:"gross_transaction_value" gorm:"type:decimal(10,2);comment:累计交易金额 GTV"`
		GrossTransactionOrder int             `json:"gross_transaction_order" gorm:"comment:累计交易订单数 GTO"`
		CurrentPointsGT       int             `json:"current_points_gt" gorm:"当前积分大于"`
		CurrentBalanceGT      decimal.Decimal `json:"current_balance_gt" gorm:"type:decimal(10,2);comment:当前余额大于"`
		ProductId             uint            `json:"product_id" gorm:"comment:商品id"`
		CreatedAt             time.Time       `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt             time.Time       `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}

	UserTag struct {
		Id        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		UserId    uint      `json:"user_id"`
		TagId     uint      `json:"tag_id" gorm:"not null;comment:会员标签id"`
		CreatedAt time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	}

	Product struct {
		Id             uint            `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		Name           string          `json:"name" gorm:"size:20;not null;comment:商品名称"`
		Img            string          `json:"img" gorm:"size:128;comment:商品图片"`
		Price          decimal.Decimal `json:"price" gorm:"type:decimal(10,2);comment:商品价格"`
		VipPrice       decimal.Decimal `json:"vip_price" gorm:"type:decimal(10,2);comment:会员价格"`
		Status         int8            `json:"status" gorm:"default=0;comment:商品状态，取值为[0:待上架, 1:出售中, 2:已下架]"`
		VipPriceStatus int8            `json:"vip_price_status" gorm:"default=0;comment:会员价状态，取值为[1:不参与, 2:参与]"`
		CreatedAt      time.Time       `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
		UpdatedAt      time.Time       `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	}
)

func (m *User) TableName() string {
	return "tb_user"
}

func (m *VipStrategy) TableName() string {
	return "tb_vip_strategy"
}

func (m *VipTag) TableName() string {
	return "tb_vip_tag"
}

func (m *UserTag) TableName() string {
	return "tb_user_tag"
}

func (m *Product) TableName() string {
	return "tb_product"
}

func createTable() error {
	err := mysql.GetDB().AutoMigrate(&User{})
	if err != nil {
		log.Printf("创建 tb_user/tb_vip_level/tb_vip_strategy/tb_vip_tag/tb_user_tag/tb_product 表失败, err: %s", err)
		return err
	}

	return nil
}
