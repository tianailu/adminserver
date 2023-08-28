package service

import (
	"errors"
	status "github.com/tianailu/adminserver/api/admin/content_mgr/common"
	"github.com/tianailu/adminserver/api/admin/content_mgr/domain"
	"github.com/tianailu/adminserver/api/admin/content_mgr/domain/req"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"gorm.io/gorm"
	"time"
)

type TagService struct {
	db *gorm.DB
}

func NewTagService() TagService {
	return TagService{
		db: mysql.GetDB(),
	}
}

func (svc *TagService) SaveTag(accountId string, tagName string) error {
	var tag domain.Tag
	err := svc.db.Model(&domain.Tag{}).Where("name = ? ", tagName).First(&tag).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("标签名已存在")
	}
	tag = domain.Tag{
		Name:            tagName,
		Status:          status.TAG_ENABLE,
		CreateAccountId: accountId,
		CreateAt:        time.Now().UnixMilli(),
	}
	return svc.db.Save(&tag).Error
}

func (svc *TagService) UpdateTag(accountId string, tagId string, tagName string) error {
	tag := domain.Tag{}
	err := svc.db.Model(domain.Tag{}).Where("id = ? ", tagId).First(&tag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("标签名不存在")
	}
	tag = domain.Tag{}
	err = svc.db.Model(domain.Tag{}).Where("name = ? ", tagName).First(&tag).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("标签名已存在")
	}
	tag.Name = tagName
	tag.UpdateAccountId = accountId
	tag.UpdateAt = time.Now().UnixMilli()
	return svc.db.Model(domain.Tag{}).Updates(tag).Error
}

func (svc *TagService) DeleteTagById(tagId string) error {
	// fixme delete user reference tag
	return svc.db.Where("id = ?", tagId).Delete(&domain.Tag{}).Error
}

func (svc *TagService) BatchDeleteTags(tagIds []int) error {
	// fixme delete user reference tag
	return svc.db.Where("id in {?}", tagIds).Delete(&domain.Tag{}).Error
}

func (svc *TagService) EnableTag(tagId int) error {
	tag := domain.Tag{}
	err := svc.db.First(&tag, tagId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("标签不存在")
	}
	return svc.db.UpdateColumn("status", status.TAG_ENABLE).Error
}

func (svc *TagService) DisableTag(tagId int) error {
	tag := domain.Tag{}
	err := svc.db.First(&tag, tagId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("标签不存在")
	}
	return svc.db.UpdateColumn("status", status.TAG_DISABLE).Error
}

func (svc *TagService) GetTagsPage(req.TagQueryReq) (common.ResponseData, error) {

	var dataResp = common.ResponseData{}
	return dataResp, nil
}
