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

const findCompanionDetailResult = `
fca.id AS find_companion_activity_id,
fcr.id AS find_companion_request_id, 
u1.user_id AS sender_user_id, 
u1.name AS sender_name, 
u1.avatar AS sender_avatar, 
u1.gender AS sender_gender, 
u1.identity_tag AS sender_identity_tag, 
u2.user_id AS receiver_user_id, 
u2.name AS receiver_name, 
u2.avatar AS receiver_avatar, 
u2.gender AS receiver_gender, 
u2.identity_tag AS receiver_identity_tag, 
c.id AS companion_type_id,
c.name AS companion_type_name,
fca.status AS matching_status, 
fca.create_at AS publishing_activity_time,
fca.cancel_time AS activity_cancel_time
fcr.created_at AS request_time,
fcr.approved_time AS sender_confirm_time
`

type FindCompanionRepo struct {
	db *gorm.DB
}

func NewFindCompanionRepo(db *gorm.DB) *FindCompanionRepo {
	return &FindCompanionRepo{
		db: db,
	}
}

func (r *FindCompanionRepo) CreateCompanionType(ctx context.Context, companionType *models.CompanionType) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(companionType).Error
	})

	if err != nil {
		r.db.Logger.Error(ctx, "Create companionType error: %s", err)
		return err
	}

	return nil
}

func (r *FindCompanionRepo) BatchUpdateCompanionType(ctx context.Context, companionTypes []*models.CompanionType) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, ct := range companionTypes {
			err := tx.Model(&models.CompanionType{}).Where("id = ?", ct.Id).Updates(ct).Error
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		r.db.Logger.Error(ctx, "Batch update companionType error: %s", err)
		return err
	}

	return nil
}

func (r *FindCompanionRepo) FindAllCompanionType(ctx context.Context, status int8) ([]*models.CompanionType, bool, error) {
	var result []*models.CompanionType

	db := r.db.WithContext(ctx).Model(&models.CompanionType{})
	if status > 0 {
		db = db.Where("status = ?", status)
	}

	err := db.Find(&result).Error

	if err != nil {
		return nil, false, err
	}

	if len(result) <= 0 {
		return result, false, nil
	}

	return result, true, nil
}

func (r *FindCompanionRepo) FindCompanionTypeByIds(ctx context.Context, ids []uint) ([]*models.CompanionType, bool, error) {
	var result []*models.CompanionType

	err := r.db.WithContext(ctx).Model(&models.CompanionType{}).Where("id IN ?", ids).Find(&result).Error

	if err != nil {
		return nil, false, err
	}

	if len(result) <= 0 {
		return result, false, nil
	}

	return result, true, nil
}

func (r *FindCompanionRepo) Find(ctx context.Context, param *models.FindCompanionSearchParam) ([]*models.FindCompanionDetail, bool, error) {
	var list []*models.FindCompanionDetail

	offset, size := page.CalPageOffset(param.PageNum, param.PageSize)

	db := r.db.WithContext(ctx).
		Select(findCompanionDetailResult).
		Table(new(models.FindCompanionActivity).TableName() + " AS fca").
		Joins("JOIN" + new(models.FindCompanionRequest).TableName() + " AS fcr ON fca.Id = fcr.activity_id").
		Joins("JOIN " + new(models.User).TableName() + " AS u1 ON fca.user_id = u1.user_id").
		Joins("JOIN " + new(models.User).TableName() + " AS u2 ON fcr.applicant_user_id = u2.user_id").
		Joins("JOIN" + new(models.CompanionType).TableName() + " AS c ON fca.companion_type_id = c.id").
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

	if param.CompanionTag > 0 {
		db = db.Where("c.tag = ?", param.CompanionTag)
	}

	if param.MatchingStatus > 0 {
		db = db.Where("fca.matching_status = ?", param.MatchingStatus)
	}

	var startAt time.Time
	var endAt time.Time
	if param.PublishStartAt > 0 {
		startAt = time.UnixMilli(param.PublishStartAt)
	}
	if param.PublishEndAt > 0 {
		endAt = time.UnixMilli(param.PublishEndAt)
	} else {
		endAt = time.Now()
	}

	if !startAt.IsZero() && !endAt.IsZero() {
		db = db.Where("fca.created_at BETWEEN ? AND ?", startAt, endAt)
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

func (r *FindCompanionRepo) TotalFindCompanion(ctx context.Context, param *models.FindCompanionSearchParam) (int64, error) {
	var count int64

	db := r.db.WithContext(ctx).
		Select(findCompanionDetailResult).
		Table(new(models.FindCompanionActivity).TableName() + " AS fca").
		Joins("JOIN" + new(models.FindCompanionRequest).TableName() + " AS fcr ON fca.Id = fcr.activity_id").
		Joins("JOIN " + new(models.User).TableName() + " AS u1 ON fca.user_id = u1.user_id").
		Joins("JOIN " + new(models.User).TableName() + " AS u2 ON fcr.applicant_user_id = u2.user_id").
		Joins("JOIN" + new(models.CompanionType).TableName() + " AS c ON fca.companion_type_id = c.id")

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

	if param.CompanionTag > 0 {
		db = db.Where("c.tag = ?", param.CompanionTag)
	}

	if param.MatchingStatus > 0 {
		db = db.Where("fca.matching_status = ?", param.MatchingStatus)
	}

	var startAt time.Time
	var endAt time.Time
	if param.PublishStartAt > 0 {
		startAt = time.UnixMilli(param.PublishStartAt)
	}
	if param.PublishEndAt > 0 {
		endAt = time.UnixMilli(param.PublishEndAt)
	} else {
		endAt = time.Now()
	}

	if !startAt.IsZero() && !endAt.IsZero() {
		db = db.Where("fca.created_at BETWEEN ? AND ?", startAt, endAt)
	}

	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
