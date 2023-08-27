package repo

import (
	"github.com/tianailu/adminserver/api/admin/user/models"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

type AboutMeRepo struct {
	db *gorm.DB
}

func NewAboutMeRepo(db *gorm.DB) *AboutMeRepo {
	return &AboutMeRepo{
		db: db,
	}
}

func (r *AboutMeRepo) Create(ctx context.Context, aboutMe *models.AboutMe) error {
	now := time.Time{}
	aboutMe.CreatedAt = now.UnixMilli()
	aboutMe.UpdatedAt = now.UnixMilli()

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(aboutMe).Error
	})

	if err != nil {
		r.db.Logger.Error(ctx, "Create aboutMe data error: %s", err)
		return err
	}

	return nil
}

func (r *AboutMeRepo) FindByUserId(ctx context.Context, userId int64) (*models.AboutMe, bool, error) {
	var am *models.AboutMe

	err := r.db.WithContext(ctx).
		Model(&models.AboutMe{}).
		Where("user_id = ?", userId).
		First(&am).Error

	if err == gorm.ErrRecordNotFound {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return am, true, nil
}
