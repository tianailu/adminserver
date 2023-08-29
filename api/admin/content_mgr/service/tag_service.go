package service

import (
	"errors"
	status "github.com/tianailu/adminserver/api/admin/content_mgr/common"
	"github.com/tianailu/adminserver/api/admin/content_mgr/domain"
	"github.com/tianailu/adminserver/api/admin/content_mgr/domain/req"
	"github.com/tianailu/adminserver/api/admin/content_mgr/domain/resp"
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
	return svc.db.Model(&domain.Tag{}).Where("id = ? ", tagId).Updates(tag).Error
}

func (svc *TagService) DeleteTagById(tagId string) error {
	// fixme delete user reference tag
	return svc.db.Where("id = ?", tagId).Delete(&domain.Tag{}).Error
}

func (svc *TagService) BatchDeleteTags(tagIds []int) error {
	// fixme delete user reference tag
	return svc.db.Where("id in ?", tagIds).Delete(&domain.Tag{}).Error
}

func (svc *TagService) EnableTag(tagId int) error {
	tag := domain.Tag{}
	err := svc.db.First(&tag, tagId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("标签不存在")
	}
	return svc.db.Model(&domain.Tag{}).Where("id = ?", tagId).UpdateColumn("status", status.TAG_ENABLE).Error
}

func (svc *TagService) DisableTag(tagId int) error {
	tag := domain.Tag{}
	err := svc.db.First(&tag, tagId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("标签不存在")
	}
	return svc.db.Model(&domain.Tag{}).Where("id = ?", tagId).UpdateColumn("status", status.TAG_DISABLE).Error
}

func (svc *TagService) GetTagsPage(reqParam req.TagQueryReq) (total int64, retData []resp.TagPageItem, error error) {
	var tags []domain.Tag
	offset := (reqParam.PageNum - 1) * reqParam.PageSize
	var cnt int64
	err := svc.db.Model(&domain.Tag{}).Where("name like ?", "%"+reqParam.Keyword+"%").Count(&cnt).Error
	if err != nil {
		return 0, nil, err
	}
	err = svc.db.Model(&domain.Tag{}).Where("name like ?", "%"+reqParam.Keyword+"%").Order("create_at desc").Limit(reqParam.PageSize).Offset(offset).Find(&tags).Error
	if err != nil {
		return 0, nil, err
	}

	var data []resp.TagPageItem
	for _, tag := range tags {
		item := resp.TagPageItem{
			Id:         tag.Id,
			Name:       tag.Name,
			Priority:   0,
			UsageCount: 0,
			CreateTime: tag.CreateAt,
			Status:     tag.Status,
		}
		data = append(data, item)
	}
	return cnt, data, err
}
