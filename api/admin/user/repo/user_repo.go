package repo

import (
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user"
	"gorm.io/gorm"
)

type UserRepo struct {
	echo.Logger
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB, log echo.Logger) *UserRepo {
	return &UserRepo{
		Logger: log,
		db:     db,
	}
}

func (r *UserRepo) Create(user *user.User) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(user).Error
	})

	if err != nil {
		r.Errorf("Create user error: %s", err)
		return err
	}

	return nil
}

func (r *UserRepo) FindByUid(uid string) (*user.User, error) {
	var u *user.User

	err := r.db.Where("uid = ?", uid).First(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepo) TotalUser() (int64, error) {
	var count int64
	err := r.db.Model(&user.User{}).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
