package services

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/api/admin/user/repo"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/utility/json"
	"github.com/tianailu/adminserver/pkg/utility/times"
)

type HeartbeatService struct {
	echo.Logger
	heartbeatRepo *repo.HeartbeatRepo
}

func NewHeartbeatService() *HeartbeatService {
	return &HeartbeatService{
		heartbeatRepo: repo.NewHeartbeatRepo(mysql.GetDB()),
	}
}

func (l *HeartbeatService) Find(ctx context.Context, param *models.HeartbeatSearchParam) ([]*models.HeartbeatListItem, int, int, int64, error) {
	var result []*models.HeartbeatListItem

	heartbeats, found, err := l.heartbeatRepo.Find(ctx, param)
	if err != nil {
		l.Errorf("Failed to find friend ship list with param, param: %s, err: %s", json.ToJsonString(param), err)
		return nil, 0, 0, 0, err
	} else if found {
		return result, param.PageNum, param.PageSize, 0, nil
	}

	totalFriend, err := l.heartbeatRepo.TotalHeartbeat(ctx, param)
	if err != nil {
		return result, 0, 0, 0, err
	}

	for _, heartbeat := range heartbeats {
		result = append(result, &models.HeartbeatListItem{
			HeartbeatRequestId:  heartbeat.HeartbeatRequestId,
			SenderUserId:        heartbeat.SenderUserId,
			SenderName:          heartbeat.SenderName,
			SenderAvatar:        heartbeat.SenderAvatar,
			SenderGender:        heartbeat.SenderGender,
			SenderIdentityTag:   heartbeat.SenderIdentityTag,
			RequestTime:         times.ToMillisecond(heartbeat.RequestTime),
			ReceiverUserId:      heartbeat.ReceiverUserId,
			ReceiverName:        heartbeat.ReceiverName,
			ReceiverAvatar:      heartbeat.ReceiverAvatar,
			ReceiverGender:      heartbeat.ReceiverGender,
			ReceiverIdentityTag: heartbeat.ReceiverIdentityTag,
			ReceiverConfirmTime: times.ToMillisecond(heartbeat.ReceiverConfirmTime),
			MatchingStatus:      heartbeat.MatchingStatus,
		})
	}

	return result, param.PageNum, param.PageSize, totalFriend, nil
}
