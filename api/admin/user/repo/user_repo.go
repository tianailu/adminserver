package repo

import (
	"context"
	"database/sql"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/pkg/utility/page"
	"gorm.io/gorm"
	"unicode/utf8"
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

func (r *UserRepo) Find(ctx context.Context, param *models.UserSearchParam) ([]*models.User, bool, error) {
	var list []*models.User

	offset, size := page.CalPageOffset(param.PageNum, param.PageSize)

	db := r.db.WithContext(ctx).Model(&models.User{}).Offset(offset).Limit(size)

	if utf8.RuneCountInString(param.Keywords) > 0 {
		db = db.Where("user_id LIKE ?", "%"+param.Keywords+"%").
			Or("name LIKE ?", "%"+param.Keywords+"%")
	}
	if param.Gender > 0 {
		db = db.Where("gender = ?", param.Gender)
	}
	if param.IdentityTag > 0 {
		db = db.Where("identity_tag = ?", param.IdentityTag)
	}
	if param.IsVip > 0 {
		db = db.Where("is_vip = ?", param.IsVip)
	}
	if param.VipTag > 0 {
		db = db.Where("vip_tag = ?", param.VipTag)
	}
	if param.AuditStatus > 0 {
		db = db.Where("vip_tag = ?", param.AuditStatus)
	}
	if param.Recommend > 0 {
		db = db.Where("recommend = ?", param.Recommend)
	}
	if utf8.RuneCountInString(param.RegisterPlace) > 0 {
		db = db.Where("register_place = ?", param.RegisterPlace)
	}
	if param.RegisterSource > 0 {
		db = db.Where("register_source = ?", param.RegisterSource)
	}

	err := db.Find(&list).Error
	if err != nil {
		return list, false, err
	}

	if len(list) <= 0 {
		return list, false, nil
	}

	return list, true, nil
}

func (r *UserRepo) FindByUserId(ctx context.Context, userId int64) (*models.User, bool, error) {
	var u *models.User

	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("user_id = ?", userId).
		First(&u).Error

	if err == gorm.ErrRecordNotFound {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return u, true, nil
}

func (r *UserRepo) TotalUser(ctx context.Context, param *models.UserSearchParam) (int64, error) {
	var count int64

	db := r.db.WithContext(ctx).Model(&models.User{})

	if utf8.RuneCountInString(param.Keywords) > 0 {
		db = db.Where("user_id LIKE ?", "%"+param.Keywords+"%").
			Or("name LIKE ?", "%"+param.Keywords+"%")
	}
	if param.Gender > 0 {
		db = db.Where("gender = ?", param.Gender)
	}
	if param.IdentityTag > 0 {
		db = db.Where("identity_tag = ?", param.IdentityTag)
	}
	if param.IsVip > 0 {
		db = db.Where("is_vip = ?", param.IsVip)
	}
	if param.VipTag > 0 {
		db = db.Where("vip_tag = ?", param.VipTag)
	}
	if param.AuditStatus > 0 {
		db = db.Where("vip_tag = ?", param.AuditStatus)
	}
	if param.Recommend > 0 {
		db = db.Where("recommend = ?", param.Recommend)
	}
	if utf8.RuneCountInString(param.RegisterPlace) > 0 {
		db = db.Where("register_place = ?", param.RegisterPlace)
	}
	if param.RegisterSource > 0 {
		db = db.Where("register_source = ?", param.RegisterSource)
	}

	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *UserRepo) MaxUserId(ctx context.Context) (int64, error) {
	var maxUserId sql.NullInt64
	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Select("MAX(user_id)").
		Find(&maxUserId).Error
	if err != nil {
		return 0, err
	}

	if maxUserId.Valid {
		return maxUserId.Int64, nil
	}

	return 0, nil
}
