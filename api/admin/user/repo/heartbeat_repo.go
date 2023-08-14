package repo

import (
	"context"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/pkg/utility/common"
	"github.com/tianailu/adminserver/pkg/utility/page"
	"gorm.io/gorm"
	"time"
	"unicode/utf8"
)

const heartbeatDetailResult = `
hr.id AS heartbeat_request_id, 
hr.sender_user_id AS sender_user_id, 
u1.name AS sender_name, 
u1.avatar AS sender_avatar, 
u1.gender AS sender_gender, 
u1.identity_tag AS sender_identity_tag, 
hr.receiver_user_id AS receiver_user_id, 
u2.name AS receiver_name, 
u2.avatar AS receiver_avatar, 
u2.gender AS receiver_gender, 
u2.identity_tag AS receiver_identity_tag, 
hr.matching_status AS matching_status, 
hr.receiver_confirm_time AS receiver_confirm_time, 
hr.created_at AS request_time 
`

type HeartbeatRepo struct {
	db *gorm.DB
}

func NewHeartbeatRepo(db *gorm.DB) *HeartbeatRepo {
	return &HeartbeatRepo{
		db: db,
	}
}

func (r *HeartbeatRepo) Find(ctx context.Context, param *models.HeartbeatSearchParam) ([]*models.HeartbeatDetail, bool, error) {
	var list []*models.HeartbeatDetail

	offset, size := page.CalPageOffset(param.PageNum, param.PageSize)

	db := r.db.WithContext(ctx).
		Select(heartbeatDetailResult).
		Table(new(models.HeartbeatRequest).TableName() + " AS hr").
		Joins("JOIN " + new(models.User).TableName() + " AS u1 ON hr.sender_user_id = u1.user_id").
		Joins("JOIN " + new(models.User).TableName() + " AS u2 ON hr.receiver_user_id = u2.user_id").
		Offset(offset).Limit(size)

	if utf8.RuneCountInString(param.Keywords) > 0 {
		if param.MatchType == 0 {
			if common.IsNumeric(param.Keywords) {
				db = db.Where("(u1.user_id = ? or u1.name LIKE ? or u2.user_id = ? or u2.name LIKE ?)",
					common.ToInt(param.Keywords), "%"+param.Keywords+"%", common.ToInt(param.Keywords), "%"+param.Keywords+"%")
			} else {
				db = db.Where("(u1.name LIKE ? or u2.name LIKE ?)", "%"+param.Keywords+"%", "%"+param.Keywords+"%")
			}
		} else if param.MatchType == 1 {
			if common.IsNumeric(param.Keywords) {
				db = db.Where("(u1.user_id = ? or u1.name LIKE ?)", common.ToInt(param.Keywords), "%"+param.Keywords+"%")
			} else {
				db = db.Where("u1.name LIKE ?", "%"+param.Keywords+"%")
			}
		} else if param.MatchType == 2 {
			if common.IsNumeric(param.Keywords) {
				db = db.Where("(u2.user_id = ? or u2.name LIKE ?)", common.ToInt(param.Keywords), "%"+param.Keywords+"%")
			} else {
				db = db.Where("u2.name LIKE ?", "%"+param.Keywords+"%")
			}
		}

		var startAt time.Time
		var endAt time.Time
		if param.ApplicationStartAt > 0 {
			startAt = time.UnixMilli(param.ApplicationStartAt)
		}
		if param.ApplicationEndAt > 0 {
			endAt = time.UnixMilli(param.ApplicationEndAt)
		} else {
			endAt = time.Now()
		}

		if !startAt.IsZero() && !endAt.IsZero() {
			db = db.Where("fr.created_at BETWEEN ? AND ?", startAt, endAt)
		}
	}

	if param.MatchingStatus > 0 {
		db = db.Where("hr.matching_status = ?", param.MatchingStatus)
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

func (r *HeartbeatRepo) TotalHeartbeat(ctx context.Context, param *models.HeartbeatSearchParam) (int64, error) {
	var count int64

	db := r.db.WithContext(ctx).
		Select(heartbeatDetailResult).
		Table(new(models.HeartbeatRequest).TableName() + " AS hr").
		Joins("JOIN " + new(models.User).TableName() + " AS u1 ON hr.sender_user_id = u1.user_id").
		Joins("JOIN " + new(models.User).TableName() + " AS u2 ON hr.receiver_user_id = u2.user_id")

	if utf8.RuneCountInString(param.Keywords) > 0 {
		if param.MatchType == 0 {
			if common.IsNumeric(param.Keywords) {
				db = db.Where("(u1.user_id = ? or u1.name LIKE ? or u2.user_id = ? or u2.name LIKE ?)",
					common.ToInt(param.Keywords), "%"+param.Keywords+"%", common.ToInt(param.Keywords), "%"+param.Keywords+"%")
			} else {
				db = db.Where("(u1.name LIKE ? or u2.name LIKE ?)", "%"+param.Keywords+"%", "%"+param.Keywords+"%")
			}
		} else if param.MatchType == 1 {
			if common.IsNumeric(param.Keywords) {
				db = db.Where("(u1.user_id = ? or u1.name LIKE ?)", common.ToInt(param.Keywords), "%"+param.Keywords+"%")
			} else {
				db = db.Where("u1.name LIKE ?", "%"+param.Keywords+"%")
			}
		} else if param.MatchType == 2 {
			if common.IsNumeric(param.Keywords) {
				db = db.Where("(u2.user_id = ? or u2.name LIKE ?)", common.ToInt(param.Keywords), "%"+param.Keywords+"%")
			} else {
				db = db.Where("u2.name LIKE ?", "%"+param.Keywords+"%")
			}
		}

		var startAt time.Time
		var endAt time.Time
		if param.ApplicationStartAt > 0 {
			startAt = time.UnixMilli(param.ApplicationStartAt)
		}
		if param.ApplicationEndAt > 0 {
			endAt = time.UnixMilli(param.ApplicationEndAt)
		} else {
			endAt = time.Now()
		}

		if !startAt.IsZero() && !endAt.IsZero() {
			db = db.Where("fr.created_at BETWEEN ? AND ?", startAt, endAt)
		}
	}

	if param.MatchingStatus > 0 {
		db = db.Where("hr.matching_status = ?", param.MatchingStatus)
	}

	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *HeartbeatRepo) CountRequestByUserIds(ctx context.Context, userIds []int64) ([]*models.HeartbeatRequestStat, error) {
	var result []*models.HeartbeatRequestStat

	err := r.db.WithContext(ctx).
		Model(&models.HeartbeatRequest{}).
		Select("sender_user_id AS user_id, count(*) AS request_count, SUM(CASE WHEN matching_status = 2 THEN 1 ELSE 0 END) AS request_success_count").
		Where("sender_user_id IN ?", userIds).
		Group("sender_user_id").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
