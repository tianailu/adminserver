package services

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/api/admin/user/repo"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/db/redis"
	pkgError "github.com/tianailu/adminserver/pkg/errors"
	"github.com/tianailu/adminserver/pkg/utility/json"
	"github.com/tianailu/adminserver/pkg/utility/times"
	"math/rand"
	"strconv"
	"time"
)

type UserService struct {
	echo.Logger
	userRepo    *repo.UserRepo
	aboutMeRepo *repo.AboutMeRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo:    repo.NewUserRepo(mysql.GetDB()),
		aboutMeRepo: repo.NewAboutMeRepo(mysql.GetDB()),
	}
}

func (l *UserService) SetLogger(logger echo.Logger) {
	l.Logger = logger
}

func (l *UserService) Find(ctx context.Context, param *models.UserSearchParam) ([]*models.UserListItem, int, int, int64, error) {
	var result []*models.UserListItem

	users, found, err := l.userRepo.Find(ctx, param)
	if err != nil {
		l.Errorf("Failed to find user list with param, param: %s, err: %s", json.ToJsonString(param), err)
		return nil, 0, 0, 0, err
	} else if !found {
		return result, param.PageNum, param.PageSize, 0, nil
	}

	totalUser, err := l.userRepo.TotalUser(ctx, param)
	if err != nil {
		return result, 0, 0, 0, err
	}

	for _, user := range users {
		result = append(result, &models.UserListItem{
			UserId:         user.UserId,
			Name:           user.Name,
			Gender:         user.Gender,
			AuditStatus:    user.AuditStatus,
			IdentityTag:    user.IdentityTag,
			IsVip:          user.IsVip,
			VipTag:         user.VipTag,
			Recommend:      user.Recommend,
			RegisterPlace:  user.RegisterPlace,
			RegisterSource: user.RegisterSource,
			RegisterTime:   user.CreatedAt.UnixMilli(),
			DurationOfUse:  user.DurationOfUse,
		})
	}

	return result, param.PageNum, param.PageSize, totalUser, nil
}

func (l *UserService) FindUserDetail(ctx context.Context, userId int64) (*models.UserDetail, error) {
	user, ok, err := l.userRepo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, pkgError.DatabaseRecordNotFound
	}

	aboutMe, ok, err := l.aboutMeRepo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	} else if !ok {
		aboutMe = &models.AboutMe{UserId: userId}
	}

	userDetail := &models.UserDetail{
		Id:               user.Id,
		AccountId:        user.AccountId,
		UserId:           user.UserId,
		Name:             user.Name,
		Avatar:           user.Avatar,
		Gender:           user.Gender,
		Birthday:         times.ToMillisecond(user.Birthday),
		Constellation:    user.Constellation,
		Height:           user.Height,
		Weight:           user.Weight,
		Education:        user.Education,
		EduStatus:        user.EduStatus,
		School:           user.School,
		Work:             user.Work,
		Company:          user.Company,
		IdentityTag:      user.IdentityTag,
		IsVip:            user.IsVip,
		VipTag:           user.VipTag,
		RegisterPlace:    user.RegisterPlace,
		RegisterSource:   user.RegisterSource,
		RegisterTime:     times.ToMillisecond(user.CreatedAt),
		AuditStatus:      user.AuditStatus,
		UserStatus:       user.UserStatus,
		TotalUsageTime:   0, // TODO 总使用时长
		Habit:            aboutMe.Habit,
		ConsumptionView:  aboutMe.ConsumptionView,
		FamilyBackground: aboutMe.FamilyBackground,
		Interest:         aboutMe.Interest,
		LoveView:         aboutMe.LoveView,
		TargetAppearance: aboutMe.TargetAppearance,
		BeImpressed:      aboutMe.BeImpressed,
	}

	return userDetail, nil
}

func (l *UserService) AddUser(ctx context.Context, userDetail *models.UserDetail) error {
	if userDetail == nil {
		return errors.New("user detail can not be empty")
	}

	newUserId, err := l.CreateUserId(ctx)
	if err != nil {
		l.Errorf("Failed to create new userId, err: %s", err)
		return err
	}

	user := &models.User{
		AccountId:      userDetail.AccountId,
		UserId:         newUserId,
		Name:           userDetail.Name,
		Avatar:         userDetail.Avatar,
		Gender:         userDetail.Gender,
		Birthday:       times.ToSqlNullTime(userDetail.Birthday),
		Constellation:  userDetail.Constellation,
		Height:         userDetail.Height,
		Weight:         userDetail.Weight,
		Education:      userDetail.Education,
		EduStatus:      userDetail.EduStatus,
		School:         userDetail.School,
		Work:           userDetail.Work,
		Company:        userDetail.Company,
		Income:         userDetail.Income,
		Residence:      userDetail.Residence,
		Hometown:       userDetail.Hometown,
		MobilePhone:    userDetail.MobilePhone,
		IdentityTag:    userDetail.IdentityTag,
		VipTag:         userDetail.VipTag,
		RegisterPlace:  userDetail.RegisterPlace,
		RegisterSource: userDetail.RegisterSource,
		DurationOfUse:  0,
	}

	if err := l.userRepo.Create(ctx, user); err != nil {
		return err
	}

	aboutMe := &models.AboutMe{
		UserId:           user.UserId,
		Habit:            userDetail.Habit,
		ConsumptionView:  userDetail.ConsumptionView,
		FamilyBackground: userDetail.FamilyBackground,
		Interest:         userDetail.Interest,
		LoveView:         userDetail.LoveView,
		TargetAppearance: userDetail.TargetAppearance,
		BeImpressed:      userDetail.BeImpressed,
	}

	if err := l.aboutMeRepo.Create(ctx, aboutMe); err != nil {
		return err
	}

	return nil
}

func (l *UserService) CreateUserId(ctx context.Context) (int64, error) {
	lockKey := redis.UserIdPollLockKey
	lock := redis.NewLock(redis.GetRDB(), lockKey, common.DefaultLockTTL, common.DefaultLockRetryInternal)
	err := lock.Lock(ctx)
	if err != nil {
		l.Errorf("Failed acquire redis lock, lockKey: %s, err: %s", lockKey, err)
		return 0, err
	}
	defer func(lock *redis.Lock, ctx context.Context) {
		err := lock.Unlock(ctx)
		if err != nil {
			l.Errorf("Failed release redis lock, lockKey: %s, err: %s", lockKey, err)
		}
	}(lock, ctx)

	newUserId := int64(-1)
	cacheKey := redis.UserIdPoolCacheKey
	userIdVal, err := redis.GetRDB().LPop(ctx, cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			curMaxUserId, err := l.userRepo.MaxUserId(ctx)
			if err != nil {
				l.Errorf("Failed to get max userId in db, err: %s", err)
				return -1, pkgError.DatabaseInternalError
			}
			if curMaxUserId < 1000 {
				curMaxUserId = 1000
			}

			var newUserIdPool []string
			for i := int64(2); i <= 10; i++ {
				newUserIdPool = append(newUserIdPool, strconv.FormatInt(curMaxUserId+i, 10))
			}
			// 随机种子
			rand.Seed(time.Now().UnixNano())
			// Fisher-Yates 算法随机打乱数组
			for i := len(newUserIdPool) - 1; i > 0; i-- {
				j := rand.Intn(i + 1)
				newUserIdPool[i], newUserIdPool[j] = newUserIdPool[j], newUserIdPool[i]
			}

			if _, err := redis.GetRDB().RPush(ctx, cacheKey, newUserIdPool).Result(); err != nil {
				l.Errorf("Failed to push userId pool item to cache, cacheKey: %s, item: %+v, err: %s", cacheKey, newUserIdPool, err)
				return -1, pkgError.RedisInternalError
			}

			newUserId = curMaxUserId + 1
		} else {
			l.Errorf("Failed to LPop new userId in cache, cacheKey: %s, err: %s", cacheKey, err)
			return -1, pkgError.RedisInternalError
		}
	} else {
		newUserId, err = strconv.ParseInt(userIdVal, 10, 64)
		if err != nil {
			l.Errorf("Incorrect userId fetched from cache, userIdVal: %s, err: %s", userIdVal, err)
			return -1, pkgError.InternalError
		}
	}

	return newUserId, nil
}