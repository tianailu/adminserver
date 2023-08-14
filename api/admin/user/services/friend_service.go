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

type FriendService struct {
	echo.Logger
	friendRepo *repo.FriendRepo
}

func NewFriendService() *FriendService {
	return &FriendService{
		friendRepo: repo.NewFriendRepo(mysql.GetDB()),
	}
}

func (l *FriendService) Find(ctx context.Context, param *models.FriendSearchParam) ([]*models.FriendListItem, int, int, int64, error) {
	var result []*models.FriendListItem

	friends, found, err := l.friendRepo.Find(ctx, param)
	if err != nil {
		l.Errorf("Failed to find friend ship list with param, param: %s, err: %s", json.ToJsonString(param), err)
		return nil, 0, 0, 0, err
	} else if !found {
		return result, param.PageNum, param.PageSize, 0, nil
	}

	totalFriend, err := l.friendRepo.TotalFriend(ctx, param)
	if err != nil {
		return result, 0, 0, 0, err
	}

	for _, friend := range friends {
		result = append(result, &models.FriendListItem{
			FriendRequestId:     friend.FriendRequestId,
			SenderUserId:        friend.SenderUserId,
			SenderName:          friend.SenderName,
			SenderAvatar:        friend.SenderAvatar,
			SenderGender:        friend.SenderGender,
			SenderIdentityTag:   friend.SenderIdentityTag,
			RequestTime:         times.ToMillisecond(friend.RequestTime),
			ReceiverUserId:      friend.ReceiverUserId,
			ReceiverName:        friend.ReceiverName,
			ReceiverAvatar:      friend.ReceiverAvatar,
			ReceiverGender:      friend.ReceiverGender,
			ReceiverIdentityTag: friend.ReceiverIdentityTag,
			ReceiverConfirmTime: times.ToMillisecond(friend.ReceiverConfirmTime),
			MatchingStatus:      friend.MatchingStatus,
		})
	}

	return result, param.PageNum, param.PageSize, totalFriend, nil
}
