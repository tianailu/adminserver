package auth

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tianailu/adminserver/config"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/utility/crypto"
	"github.com/tianailu/adminserver/pkg/utility/json"
	"github.com/tianailu/adminserver/pkg/utility/page"
	"github.com/tianailu/adminserver/pkg/utility/snowflake"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Account struct {
	Id          uint                  `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:主键"`
	AccountId   string                `json:"account_id" gorm:"size:32;not null;comment:账号Id;index:idx_account_id"`
	OpenId      string                `json:"open_id" gorm:"size:32;comment:微信用户唯一标识"`
	UnionId     string                `json:"union_id" gorm:"size:32;comment:微信用户在开放平台的唯一标识符"`
	MobilePhone string                `json:"mobile_phone" gorm:"size:11;comment:绑定登录的手机号码"`
	Account     string                `json:"account" gorm:"size:32;comment:账号;index:idx_account"`
	Password    string                `json:"password" gorm:"size:64;comment:密码"`
	AccountType string                `json:"account_type" gorm:"size:10,not null;comment:账号类型，取值为[APP:小程序, ADMIN:管理后台]"`
	Name        string                `json:"name" gorm:"size:12;comment:昵称"`
	Avatar      string                `json:"avatar" gorm:"size:128;comment:头像地址"`
	Status      int8                  `json:"status" gorm:"not null;default:0;comment:账号状态，取值为[0:正常激活, 1:禁用]"`
	LoginCount  uint                  `json:"login_count" gorm:"default:0;comment:登录次数"`
	LastLoginIp string                `json:"last_login_ip" gorm:"size:20;comment:最近一次登录时的ip地址"`
	LastLoginAt sql.NullTime          `json:"last_login_at" gorm:"type:datetime;comment:最近一次登录时间"`
	CreatedAt   time.Time             `json:"created_at" gorm:"type:datetime;autoCreateTime;default:CURRENT_TIMESTAMP;not null;comment:注册时间"`
	UpdatedAt   time.Time             `json:"updated_at" gorm:"type:datetime;autoUpdateTime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;comment:修改时间"`
	DeletedAt   sql.NullTime          `json:"deleted_at" gorm:"type:datetime;comment:删除时间"`
	IsDel       soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag;comment:删除标志，取值为[0:使用中, 1:已删除]"`
	Remark      string                `json:"remark" gorm:"size:64;comment:备注"`
}

func (m *Account) TableName() string {
	return "tb_account"
}

func (m *Account) String() string {
	return json.ToJsonString(m)
}

func createTable() error {
	err := mysql.GetDB().Set("gorm:tb_account", "ENGINE=InnoDB").AutoMigrate(&Account{})
	if err != nil {
		log.Errorf("创建 tb_account 表失败, err: %s", err)
		return err
	}

	var account *Account
	err = mysql.GetDB().
		Where("account = ?", config.AuthConf.DefaultAdminAccount).
		Where("account_type = ?", DefaultAccountType).
		First(&account).Error
	if err == gorm.ErrRecordNotFound {
		account = &Account{
			AccountId:   snowflake.GetNode().Generate().String(),
			Account:     config.AuthConf.DefaultAdminAccount,
			Password:    crypto.GetSha256String(config.AuthConf.DefaultAdminPassword, config.AuthConf.AdminPasswordSalt),
			AccountType: DefaultAccountType,
			Name:        config.AuthConf.DefaultAdminName,
			Status:      0,
			DeletedAt:   sql.NullTime{Valid: false},
		}
		if err = mysql.GetDB().Create(account).Error; err != nil {
			log.Errorf("创建系统默认管理员账号失败, err: %s", err)
			return err
		}
	} else if err != nil {
		log.Errorf("Failed to query system admin account, account: %s, err: %s", config.AuthConf.DefaultAdminAccount, err)
		return err
	}

	return nil
}

type AccountRepo struct {
	echo.Logger
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB, log echo.Logger) *AccountRepo {
	return &AccountRepo{
		Logger: log,
		db:     db,
	}
}

func (r *AccountRepo) Create(account *Account) error {
	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Create(account).Error
	})

	if err != nil {
		r.Errorf("Create account error: %s", err)
		return err
	}

	return nil
}

func (r *AccountRepo) Find(name string, pageNum, pageSize int) ([]*Account, error) {
	var list []*Account

	offset, size := page.CalPageOffset(pageNum, pageSize)
	db := r.db.Offset(offset).Limit(size).Where("account_type = ?", "ADMIN")

	if len(name) > 0 {
		db = db.Where("name = ?", name)
	}

	err := db.Find(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
}

func (r *AccountRepo) FindByAccount(account, accountType string) (*Account, error) {
	var a *Account

	err := r.db.Where("account = ? AND account_type = ?", account, accountType).First(&a).Error

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (r *AccountRepo) FindByAccountId(accountId string) (*Account, error) {
	var a *Account

	err := r.db.Where("account_id = ?", accountId).First(&a).Error

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (r *AccountRepo) Update(account *Account) error {
	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Model(&Account{}).Where("id = ?", account.Id).Updates(account).Error
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepo) UpdateStatus(userId, status string) error {
	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Model(&Account{}).Where("user_id = ?", userId).Update("status", status).Error
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepo) Login(accountId, lastLoginIp string, loginCount uint) error {
	err := mysql.GetDB().Transaction(func(tx *gorm.DB) error {
		return tx.Model(&Account{}).Where("account_id = ?", accountId).Updates(Account{
			LoginCount:  loginCount,
			LastLoginIp: lastLoginIp,
			LastLoginAt: sql.NullTime{Time: time.Now(), Valid: true},
		}).Error
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepo) TotalAdmin() (int64, error) {
	var count int64
	err := r.db.Model(&Account{}).Where("account_type = ?", "ADMIN").Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
