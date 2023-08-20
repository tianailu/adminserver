package services

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/api/admin/user/repo"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/utility/times"
)

type FindCompanionService struct {
	echo.Logger
	findCompanionRepo *repo.FindCompanionRepo
}

func NewFindCompanionService() *FindCompanionService {
	return &FindCompanionService{
		findCompanionRepo: repo.NewFindCompanionRepo(mysql.GetDB()),
	}
}

func (l *FindCompanionService) Find(ctx context.Context, param *models.FindCompanionSearchParam) ([]*models.FindCompanionListItem, int, int, int64, error) {
	var result []*models.FindCompanionListItem

	findCompanions, found, err := l.findCompanionRepo.Find(ctx, param)
	if err != nil {
		return nil, 0, 0, 0, err
	} else if !found {
		return result, param.PageNum, param.PageSize, 0, nil
	}

	totalFindCompanion, err := l.findCompanionRepo.TotalFindCompanion(ctx, param)
	if err != nil {
		return result, 0, 0, 0, err
	}

	for _, findCompanion := range findCompanions {
		result = append(result, &models.FindCompanionListItem{
			FindCompanionActivityId: findCompanion.FindCompanionActivityId,
			SenderUserId:            findCompanion.SenderUserId,
			SenderName:              findCompanion.SenderName,
			SenderAvatar:            findCompanion.SenderAvatar,
			SenderGender:            findCompanion.SenderGender,
			SenderIdentityTag:       findCompanion.SenderIdentityTag,
			ReceiverUserId:          findCompanion.ReceiverUserId,
			ReceiverName:            findCompanion.ReceiverName,
			ReceiverAvatar:          findCompanion.ReceiverAvatar,
			ReceiverGender:          findCompanion.ReceiverGender,
			ReceiverIdentityTag:     findCompanion.ReceiverIdentityTag,
			CompanionTypeId:         findCompanion.CompanionTypeId,
			CompanionTypeName:       findCompanion.CompanionTypeName,
			MatchingStatus:          findCompanion.MatchingStatus,
			PublishingActivityTime:  times.ToMillisecond(findCompanion.PublishingActivityTime),
			RequestTime:             times.ToMillisecond(findCompanion.RequestTime),
			ActivityCancelTime:      times.ToMillisecond(findCompanion.ActivityCancelTime),
			SenderConfirmTime:       times.ToMillisecond(findCompanion.SenderConfirmTime),
		})
	}

	return result, param.PageNum, param.PageSize, totalFindCompanion, nil
}

func (l *FindCompanionService) FindCompanionTypeList(ctx context.Context, param *models.CompanionTypeSearchParam) ([]*models.CompanionTypeListItem, error) {
	var result []*models.CompanionTypeListItem

	status := param.Status
	if status < 0 {
		status = 0
	}

	companionTypes, _, err := l.findCompanionRepo.FindAllCompanionType(ctx, param.Status)
	if err != nil {
		return nil, err
	}

	for _, companionType := range companionTypes {
		result = append(result, &models.CompanionTypeListItem{
			CompanionTypeId:   companionType.Id,
			CompanionTag:      companionType.Tag,
			CompanionTypeName: companionType.Name,
			Status:            companionType.Status,
			CreateAt:          times.ToMillisecond(companionType.CreatedAt),
		})
	}

	return result, nil
}

func (l *FindCompanionService) AddCompanionType(ctx context.Context, companionTag int8, companionTypeName string, status int8) error {
	if status <= 0 {
		status = 2
	}

	err := l.findCompanionRepo.CreateCompanionType(ctx, &models.CompanionType{
		Tag:    companionTag,
		Name:   companionTypeName,
		Status: status,
	})
	if err != nil {
		return err
	}

	return nil
}
