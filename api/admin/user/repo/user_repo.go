package repo

import (
	"context"
	"database/sql"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, user *models.User) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(user).Error
	})

	if err != nil {
		r.db.Logger.Error(ctx, "Create user error: %s", err)
		return err
	}

	return nil
}

func (r *UserRepo) FindByUid(ctx context.Context, uid int64) (*models.User, bool, error) {
	var u *models.User

	err := r.db.WithContext(ctx).Where("uid = ?", uid).First(&u).Error

	if err == gorm.ErrRecordNotFound {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return u, true, nil
}

func (r *UserRepo) TotalUser(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.User{}).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *UserRepo) MaxUid(ctx context.Context) (int64, error) {
	var maxUid sql.NullInt64
	err := r.db.WithContext(ctx).Model(&models.User{}).Select("MAX(uid)").Find(&maxUid).Error
	if err != nil {
		return 0, err
	}

	if maxUid.Valid {
		return maxUid.Int64, nil
	}

	return 0, nil
}
