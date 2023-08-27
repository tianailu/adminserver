package services

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/api/admin/user/repo"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/utility/times"
	"unicode/utf8"
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
			CreateAt:          companionType.CreatedAt,
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

func (l *FindCompanionService) UpdateCompanionTypeList(ctx context.Context, companionTypes []*models.CompanionTypeListItem) error {
	if len(companionTypes) == 0 {
		return nil
	}

	ids := make([]int64, 0)
	idToCT := make(map[int64]*models.CompanionTypeListItem)
	for _, ct := range companionTypes {
		ids = append(ids, ct.CompanionTypeId)
		idToCT[ct.CompanionTypeId] = ct
	}

	sourceCompanionTypes, found, err := l.findCompanionRepo.FindCompanionTypeByIds(ctx, ids)
	if err != nil {
		return err
	} else if !found {
		return nil
	}

	for _, ct := range sourceCompanionTypes {
		reqCT := idToCT[ct.Id]
		if reqCT == nil {
			continue
		}

		if reqCT.Status > 0 && reqCT.Status <= 2 {
			ct.Status = reqCT.Status
		}

		if utf8.RuneCountInString(reqCT.CompanionTypeName) > 0 {
			ct.Name = reqCT.CompanionTypeName
		}
	}

	err = l.findCompanionRepo.BatchUpdateCompanionType(ctx, sourceCompanionTypes)
	if err != nil {
		return err
	}

	return nil
}
