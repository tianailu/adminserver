package repo

import (
	"context"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"gorm.io/gorm"
)

type FollowRepo struct {
	db *gorm.DB
}

func NewFollowRepo(db *gorm.DB) *FollowRepo {
	return &FollowRepo{
		db: db,
	}
}

func (r *FollowRepo) CountFollowByUserIds(ctx context.Context, userIds []int64) ([]*models.UserFollowStat, error) {
	var result []*models.UserFollowStat

	err := r.db.WithContext(ctx).
		Model(&models.Follow{}).
		Select("user_id, count(*) as follow_count").
		Where("user_id IN ?", userIds).
		Group("user_id").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *FollowRepo) CountFollowByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(&models.Follow{}).
		Where("user_id = ?", userId).
		Find(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *FollowRepo) CountFansByUserIds(ctx context.Context, userIds []int64) ([]*models.UserFansStat, error) {
	var result []*models.UserFansStat

	err := r.db.WithContext(ctx).
		Model(&models.Fans{}).
		Select("user_id, count(*) as follow_count").
		Where("user_id IN ?", userIds).
		Group("user_id").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *FollowRepo) CountFansByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(&models.Fans{}).
		Where("user_id = ?", userId).
		Find(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *FollowRepo) CountFollowFans(ctx context.Context, userIds []int64) ([]*models.UserFollowFansStat, error) {
	var result []*models.UserFollowFansStat

	err := r.db.WithContext(ctx).
		Select("user.user_id AS user_id, COUNT(DISTINCT follow.id) AS follow_count, COUNT(DISTINCT fans.id) AS fans_count").
		Table(new(models.User).TableName()+" AS user").
		Joins("LEFT JOIN "+new(models.Follow).TableName()+" AS follow ON user.user_id = follow.user_id").
		Joins("LEFT JOIN "+new(models.Fans).TableName()+" AS fans ON user.user_id = fans.user_id").
		Where("user.user_id IN ?", userIds).
		Group("user.user_id").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
