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

const friendDetailResult = `
fr.id AS friend_request_id, 
fr.sender_user_id AS sender_user_id, 
u1.name AS sender_name, 
u1.avatar AS sender_avatar, 
u1.gender AS sender_gender, 
u1.identity_tag AS sender_identity_tag, 
fr.receiver_user_id AS receiver_user_id, 
u2.name AS receiver_name, 
u2.avatar AS receiver_avatar, 
u2.gender AS receiver_gender, 
u2.identity_tag AS receiver_identity_tag, 
fr.matching_status AS matching_status, 
fr.receiver_confirm_time AS receiver_confirm_time, 
fr.created_at AS request_time 
`

type FriendRepo struct {
	db *gorm.DB
}

func NewFriendRepo(db *gorm.DB) *FriendRepo {
	return &FriendRepo{
		db: db,
	}
}

func (r *FriendRepo) Find(ctx context.Context, param *models.FriendSearchParam) ([]*models.FriendDetail, bool, error) {
	var list []*models.FriendDetail

	offset, size := page.CalPageOffset(param.PageNum, param.PageSize)

	db := r.db.WithContext(ctx).Table("tb_friend_request AS fr").Select(friendDetailResult).
		Joins("JOIN tb_user AS u1 ON fr.sender_user_id = u1.user_id").
		Joins("JOIN tb_user AS u2 ON fr.receiver_user_id = u2.user_id").
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
	}

	if param.MatchingStatus > 0 {
		db = db.Where("fr.matching_status = ?", param.MatchingStatus)
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

	err := db.Find(&list).Error
	if err != nil {
		return list, false, err
	}

	if len(list) <= 0 {
		return list, false, nil
	}

	return list, true, nil
}

func (r *FriendRepo) TotalFriend(ctx context.Context, param *models.FriendSearchParam) (int64, error) {
	var count int64

	db := r.db.WithContext(ctx).Table("tb_friend_request AS fr").Select(friendDetailResult).
		Joins("JOIN tb_user AS u1 ON fr.sender_user_id = u1.user_id").
		Joins("JOIN tb_user AS u2 ON fr.receiver_user_id = u2.user_id")

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
	}

	if param.MatchingStatus > 0 {
		db = db.Where("fr.matching_status = ?", param.MatchingStatus)
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

	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
