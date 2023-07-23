package user

import (
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type (
	User struct {
		Id               int     `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		AccountId        string  `json:"account_id" gorm:"size:32;comment:账号ID"`
		Uid              string  `json:"uid" gorm:"size:32;not null;comment:用户ID"`
		Name             string  `json:"name" gorm:"size:12;comment:用户名"`
		Img              string  `json:"img" gorm:"size:128;comment:个人头像"`
		Gender           int8    `json:"gender" gorm:"not null;default:0;comment:性别，取值为[0:男, 1:女]"`
		Constellation    string  `json:"constellation" gorm:"size:12;comment:星座"`
		Birthday         int64   `json:"birthday" gorm:"comment:出生时间，时间戳，单位毫秒"`
		Height           float32 `json:"height" gorm:"comment:身高，单位cm"`
		Weight           float32 `json:"weight" gorm:"comment:体重，单位kg"`
		Hometown         string  `json:"hometown" gorm:"size:12;comment:家乡"`
		Education        string  `json:"education" gorm:"size:12;comment:最高学历"`
		School           string  `json:"school" gorm:"size:32;comment:毕业院校"`
		Work             string  `json:"work" gorm:"size:20;comment:职业"`
		Company          string  `json:"company" gorm:"size:20;comment:公司"`
		Income           string  `json:"income" gorm:"size:20;comment:收入"`
		MobilePhone      string  `json:"m_phone" gorm:"size:12;comment:手机号码"`
		WeChat           string  `json:"wechat" gorm:"size:20;comment:微信号"`
		Marriage         string  `json:"marriage" gorm:"size:12;comment:婚姻状态"`
		Habit            string  `json:"habit" gorm:"size:64;comment:生活习惯"`
		ConsumptionView  string  `json:"consumption_view" gorm:"size:64;comment:消费观"`
		Family           string  `json:"family" gorm:"size:64;comment:家庭背景"`
		Interest         string  `json:"interest" gorm:"size:64;comment:兴趣爱好"`
		LoveView         string  `json:"love_view" gorm:"size:64;comment:爱情观"`
		TargetAppearance string  `json:"ta_appearance" gorm:"size:64;comment:希望另一半的样子"`
		BeImpressed      string  `json:"be_impressed" gorm:"size:64;comment:对方什么最能打动自己"`
		TargetAge        string  `json:"ta_age" gorm:"size:12;comment:希望另一半身高范围"`
		TargetHeight     string  `json:"ta_height" gorm:"size:12;comment:希望另一半身高范围"`
		TargetCity       string  `json:"ta_city" gorm:"size:12;comment:希望另一半所在城市"`
		TargetHometown   string  `json:"ta_hometown" gorm:"size:12;comment:希望另一半的家乡"`
		TargetEducation  string  `json:"ta_education" gorm:"size:12;comment:希望另一半最低学历"`
		TargetMarriage   string  `json:"ta_marriage" gorm:"size:12;comment:希望另一半婚姻状态"`
		VipLevel         int8    `json:"vip_level" gorm:"comment:会员等级"`
		IdCard           string  `json:"id_card" gorm:"size:18;comment:身份证"`
		RealName         string  `json:"real_name" gorm:"size:20;comment:真实姓名"`
		IsRealNameAuth   int8    `json:"is_rn_auth" gorm:"default:0;comment:是否完成实名认证，0:未认证，1:已通过认证"`
		WorkAuthImg      string  `json:"work_auth_img" gorm:"size:128;comment:工作认证图片"`
		IsWorkAuth       int8    `json:"is_work_auth" gorm:"default:0;comment:是否完成工作认证，0:未认证，1:已通过认证"`
		EduAuthImg       string  `json:"edu_auth_img" gorm:"size:128;comment:学历认证图片"`
		IsEduAuth        int8    `json:"is_edu_auth" gorm:"default:0;comment:是否完成学历认证，0:未认证，1:已通过认证"`
		Status           int8    `json:"status" gorm:"not null;default:0;comment:未审核，1:已审核，2:vip"`
		RegisterPlace    string  `json:"register_place" gorm:"size:12;comment:注册地"`
		RegisterSource   int8    `json:"register_source" gorm:"comment:注册来源，取值为[0:未知, 1:APP, 2:小程序, 3:群组, 4:二维码, 5:管理后台]"`
		CreatedAt        int64   `json:"created_at" gorm:"comment:注册时间，单位毫秒"`
		UpdateAt         int64   `json:"update_at" gorm:"comment:修改时间，单位毫秒"`
	}

	Hometown struct {
		Country    string `json:"country"`
		Province   string `json:"province"`
		City       string `json:"city"`
		AddrDetail string `json:"addr_detail"`
	}

	VipLevel struct {
		Id                int     `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		Level             string  `json:"level" gorm:"size:10;comment:等级"`
		Name              string  `json:"name" gorm:"size:10;comment:等级名称"`
		Desc              string  `json:"desc" gorm:"size:30;comment:描述"`
		Price             float64 `json:"price" gorm:"comment:价格"`
		Discount          float32 `json:"discount" gorm:"comment:折扣"`
		DowngradeStrategy int8    `json:"downgrade_strategy" gorm:"comment:降级策略"`
		Status            int8    `json:"status" gorm:"comment:状态，取值为[0:禁用, 1:启用]"`
		CreatedAt         int64   `json:"created_at" gorm:"comment:创建时间"`
	}

	VipTag struct {
		Id        int    `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
		Name      string `json:"name" gorm:"size:10;comment:标签名称"`
		CreatedAt int64  `json:"created_at" gorm:"comment:创建时间"`
	}
)

func createTable() error {
	time.Sleep(time.Second * 5)

	err := mysql.GetDB().AutoMigrate(&User{}, &VipLevel{}, &VipTag{})
	if err != nil {
		log.Printf("创建 tb_user/tb_vip_level/tb_vip_tag 表失败, err: %s", err)
		return err
	}

	return nil
}

func (m *User) TableName() string {
	return "tb_user"
}

func (m *VipLevel) TableName() string {
	return "tb_vip_level"
}

func (m *VipTag) TableName() string {
	return "tb_vip_tag"
}

func (m *User) Create() error {
	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Create(m).Error
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *User) FindByUid(uid string) (*User, error) {
	var u *User

	err := mysql.GetDB().Where("uid = ?", uid).First(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (m *User) TotalUser() (int64, error) {
	var count int64
	err := mysql.GetDB().Model(&User{}).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
