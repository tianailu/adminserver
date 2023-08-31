package services

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/api/admin/user/repo"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/errors"
	"github.com/tianailu/adminserver/pkg/utility/json"
	"unicode/utf8"
)

type VipService struct {
	echo.Logger
	vipTagRepo       *repo.VipTagRepo
	userRepo         *repo.UserRepo
	vipTagUpdateChan chan struct{}
}

func NewVipService() *VipService {
	return &VipService{
		vipTagRepo:       repo.NewVipTagRepo(mysql.GetDB()),
		userRepo:         repo.NewUserRepo(mysql.GetDB()),
		vipTagUpdateChan: make(chan struct{}),
	}
}

func (l *VipService) AddVipTag(ctx context.Context, vipTagDetail *models.VipTagDetail) (int32, error) {
	if vipTagDetail == nil {
		return 0, errors.ParamError
	}

	vipTag := &models.VipTag{
		Name:                vipTagDetail.Name,
		TotalRechargeAmount: vipTagDetail.TotalRechargeAmount,
	}

	err := l.vipTagRepo.Create(ctx, vipTag)
	if err != nil {
		return 0, err
	}

	// TODO 需要全量更新用户vip标签，可以使用异步处理

	return vipTag.Id, nil
}

func (l *VipService) Find(ctx context.Context, param *common.SearchParam) ([]*models.VipTagDetail, int, int, int64, error) {
	var result []*models.VipTagDetail

	vipTags, found, err := l.vipTagRepo.Find(ctx, param)
	if err != nil {
		l.Errorf("Failed to find vip_tag with param, param: %s, error: %s", json.ToJsonString(param), err)
		return nil, 0, 0, 0, err
	} else if !found {
		return result, param.PageNum, param.PageSize, 0, nil
	}

	vipTagIds := make([]int32, 0)
	for _, v := range vipTags {
		vipTagIds = append(vipTagIds, v.Id)
	}

	countVips, err := l.vipTagRepo.CountVip(ctx, vipTagIds)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	totalVipTag, err := l.vipTagRepo.Total(ctx)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	for _, v := range vipTags {
		result = append(result, &models.VipTagDetail{
			Id:                  v.Id,
			Name:                v.Name,
			TotalRechargeAmount: v.TotalRechargeAmount,
			VipCount:            countVips[v.Id],
			CreateAt:            v.CreatedAt,
			UpdateAt:            v.UpdatedAt,
		})
	}

	return result, param.PageNum, param.PageSize, totalVipTag, nil
}

func (l *VipService) UpdateVipTag(ctx context.Context, vipTag *models.VipTagDetail) error {
	if vipTag == nil {
		return nil
	}

	targetVipTag, found, err := l.vipTagRepo.FindById(ctx, vipTag.Id)
	if err != nil {
		return err
	} else if !found {
		return nil
	}

	isUpdate := false
	isUpdateTotalRechargeAmount := false
	if utf8.RuneCountInString(vipTag.Name) > 0 && targetVipTag.Name != vipTag.Name {
		targetVipTag.Name = vipTag.Name
		isUpdate = true
	}
	if !targetVipTag.TotalRechargeAmount.Equal(vipTag.TotalRechargeAmount) {
		targetVipTag.TotalRechargeAmount = vipTag.TotalRechargeAmount
		isUpdate = true
		isUpdateTotalRechargeAmount = true
	}

	if isUpdate {
		if err := l.vipTagRepo.Update(ctx, targetVipTag); err != nil {
			return err
		}
	}

	if isUpdateTotalRechargeAmount {
		// TODO 需要全量更新用户vip标签，可以使用异步处理
		if err := l.ResetUserVipTag(ctx, []int32{targetVipTag.Id}); err != nil {
			return err
		}
	}

	return nil
}

func (l *VipService) BatchDeleteVipTag(ctx context.Context, vipTagIds []int32) error {
	vipTags, found, err := l.vipTagRepo.FindByIds(ctx, vipTagIds)
	if err != nil {
		return err
	} else if !found {
		return nil
	}

	vipTagIds = make([]int32, 0)
	for _, v := range vipTags {
		vipTagIds = append(vipTagIds, v.Id)
	}

	if err := l.vipTagRepo.Delete(ctx, vipTagIds); err != nil {
		return err
	}

	if err := l.ResetUserVipTag(ctx, vipTagIds); err != nil {
		return err
	}

	return nil
}

func (l *VipService) ResetUserVipTag(ctx context.Context, vipTagIds []int32) error {
	if len(vipTagIds) <= 0 {
		return nil
	}

	users, found, err := l.userRepo.FindUserByVipTag(ctx, vipTagIds)
	if err != nil {
		return err
	} else if !found {
		return nil
	}

	userIds := make([]int64, 0)
	for _, u := range users {
		userIds = append(userIds, u.UserId)
	}

	return l.ResetUserVipTagWithUsers(ctx, userIds)
}

func (l *VipService) ResetUserVipTagWithUsers(ctx context.Context, userIds []int64) error {
	// TODO 重新计算用户的vip标识
	return nil
}
