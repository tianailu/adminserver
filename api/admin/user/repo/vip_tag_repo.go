package repo

import (
	"context"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/utility/page"
	"gorm.io/gorm"
	"time"
)

type VipTagRepo struct {
	db *gorm.DB
}

func NewVipTagRepo(db *gorm.DB) *VipTagRepo {
	return &VipTagRepo{
		db: db,
	}
}

func (r *VipTagRepo) Create(ctx context.Context, vipTag *models.VipTag) error {
	now := time.Now()
	vipTag.CreatedAt = now.UnixMilli()
	vipTag.UpdatedAt = now.UnixMilli()

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(vipTag).Error
	})

	if err != nil {
		r.db.Logger.Error(ctx, "Create vip tag error: %s", err)
		return err
	}

	return nil
}

func (r *VipTagRepo) Find(ctx context.Context, param *common.SearchParam) ([]*models.VipTag, bool, error) {
	var list []*models.VipTag

	offset, size := page.CalPageOffset(param.PageNum, param.PageSize)

	err := r.db.WithContext(ctx).Model(&models.VipTag{}).Offset(offset).Limit(size).Find(&list).Error

	if err != nil {
		return list, false, err
	}

	if len(list) <= 0 {
		return list, false, nil
	}

	return list, true, nil
}

func (r *VipTagRepo) FindById(ctx context.Context, id int32) (*models.VipTag, bool, error) {
	var result *models.VipTag

	err := r.db.WithContext(ctx).Model(&models.VipTag{}).Where("id = ?", id).First(&result).Error

	if err == gorm.ErrRecordNotFound {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return result, true, nil
}

func (r *VipTagRepo) FindByIds(ctx context.Context, ids []int32) ([]*models.VipTag, bool, error) {
	var list []*models.VipTag

	if len(ids) <= 0 {
		return list, false, nil
	}

	err := r.db.WithContext(ctx).Model(&models.VipTag{}).Where("id in ?", ids).Find(&list).Error

	if err != nil {
		return list, false, err
	}

	if len(list) <= 0 {
		return list, false, nil
	}

	return list, true, nil
}

func (r *VipTagRepo) Delete(ctx context.Context, ids []int32) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&models.VipTag{}, ids).Error
	})

	if err != nil {
		r.db.Logger.Error(ctx, "Batch delete vip_tag error: %s", err)
		return err
	}

	return nil
}

func (r *VipTagRepo) Update(ctx context.Context, vipTag *models.VipTag) error {
	if vipTag == nil {
		return nil
	}

	vipTag.UpdatedAt = time.Now().UnixMilli()

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(&models.VipTag{}).Where("id = ?", vipTag.Id).Updates(vipTag).Error
	})

	if err != nil {
		r.db.Logger.Error(ctx, "Update vip_tag error: %s", err)
		return err
	}

	return nil
}

func (r *VipTagRepo) CountVip(ctx context.Context, vipTagIds []int32) (map[int32]int64, error) {
	result := make(map[int32]int64)

	err := r.db.Model(&models.User{}).
		Select("vip_tag, count(user_id)").
		Where("vip_tag in ?", vipTagIds).
		Group("vip_tag").
		Find(&result).Error

	if err != nil {
		r.db.Logger.Error(ctx, "Failed to count the number of vips, error: %s", err)
		return result, err
	}

	return result, nil
}

func (r *VipTagRepo) Total(ctx context.Context) (int64, error) {
	var count int64

	err := r.db.WithContext(ctx).Model(&models.VipTag{}).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
