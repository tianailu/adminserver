package repo

import (
	"context"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"gorm.io/gorm"
)

type FindCompanionRepo struct {
	db *gorm.DB
}

func NewFindCompanionRepo(db *gorm.DB) *FindCompanionRepo {
	return &FindCompanionRepo{
		db: db,
	}
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
