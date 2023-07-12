package auth

import (
	"database/sql"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Account struct {
	Id          uint                  `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
	UserId      string                `json:"user_id" gorm:"size:32;comment:用户id;index:idx_user_id"`
	OpenId      string                `json:"open_id" gorm:"size:32;comment:微信用户唯一标识"`
	UnionId     string                `json:"union_id" gorm:"size:32;comment:微信用户在开放平台的唯一标识符"`
	MobilePhone string                `json:"mobile_phone" gorm:"size:20;comment:绑定登录的手机号码"`
	Account     string                `json:"account" gorm:"size:32;comment:用户名;index:idx_account"`
	Password    string                `json:"password" gorm:"size:32;comment:密码"`
	AccountType string                `json:"account_type" gorm:"size:12,not null;comment:账号类型，取值为[APP:小程序账号, ADMIN:管理后台]"`
	Name        string                `json:"name" gorm:"size:12;comment:昵称"`
	Role        string                `json:"role" gorm:"size:32;comment:用户角色"`
	Avatar      string                `json:"avatar" gorm:"size:128;comment:头像地址"`
	Status      int8                  `json:"status" gorm:"not null;default:0;comment:账号状态，取值为[0:正常激活, 1:禁用]"`
	LastLoginAt sql.NullTime          `json:"last_login_at" gorm:"type:datetime;comment:最近一次登录时间"`
	CreatedAt   time.Time             `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:创建时间"`
	UpdatedAt   time.Time             `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	DeletedAt   time.Time             `json:"deleted_at" gorm:"type:datetime;comment:删除时间"`
	IsDel       soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag;comment:删除标志，取值为[0:使用中, 1:已删除]"`
	Remark      string                `json:"remark" gorm:"size:64;comment:备注"`
}

func createDatabase() error {
	time.Sleep(time.Second * 5)

	err := mysql.GetDB().Set("gorm:tb_account", "ENGINE=InnoDB").AutoMigrate(&Account{})
	if err != nil {
		return err
	}

	return nil
}

func (m *Account) TableName() string {
	return "tb_account"
}

func (m *Account) Create() error {
	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Create(m).Error
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *Account) FindByAccount(account, accountType string) (*Account, error) {
	var a *Account

	err := mysql.GetDB().Where("account = ? AND account_type = ?", account, accountType).First(a).Error

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (m *Account) FindByUserId(userId string) (*Account, error) {
	var a *Account

	err := mysql.GetDB().Where("user_id = ?", userId).First(a).Error

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (m *Account) UpdateStatus(userId, status string) error {
	err := mysql.GetDB().Where("user_id = ?", userId).Update("status", status).Error

	if err != nil {
		return err
	}

	return nil
}

func (m *Account) Login(userId string) error {
	err := mysql.GetDB().Where("user_id = ?", userId).Update("last_login_at", sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
