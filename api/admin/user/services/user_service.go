package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/api/admin/user/repo"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/db/redis"
	pkgError "github.com/tianailu/adminserver/pkg/errors"
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

func (l *UserService) FindUserDetail(ctx context.Context, uid int64) (*models.UserDetail, error) {
	user, ok, err := l.userRepo.FindByUid(ctx, uid)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, pkgError.DatabaseRecordNotFound
	}

	aboutMe, ok, err := l.aboutMeRepo.FindByUid(ctx, uid)
	if err != nil {
		return nil, err
	} else if !ok {
		aboutMe = &models.AboutMe{UserId: uid}
	}

	userDetail := &models.UserDetail{
		Id:               user.Id,
		AccountId:        user.AccountId,
		Uid:              user.Uid,
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

	newUid, err := l.CreateUid(ctx)
	if err != nil {
		l.Errorf("Failed to create new uid, err: %s", err)
		return err
	}

	user := &models.User{
		AccountId:      userDetail.AccountId,
		Uid:            newUid,
		Name:           userDetail.Name,
		Avatar:         userDetail.Avatar,
		Gender:         userDetail.Gender,
		Birthday:       sql.NullTime{Valid: false, Time: time.UnixMilli(userDetail.Birthday)},
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
		UserId:           user.Uid,
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

func (l *UserService) CreateUid(ctx context.Context) (int64, error) {
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

	newUid := int64(-1)
	cacheKey := redis.UserIdPoolCacheKey
	uidVal, err := redis.GetRDB().LPop(ctx, cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			curMaxUid, err := l.userRepo.MaxUid(ctx)
			if err != nil {
				l.Errorf("Failed to get max uid in db, err: %s", err)
				return -1, pkgError.DatabaseInternalError
			}
			if curMaxUid < 1000 {
				curMaxUid = 1000
			}

			var newUidPool []string
			for i := int64(2); i <= 10; i++ {
				newUidPool = append(newUidPool, strconv.FormatInt(curMaxUid+i, 10))
			}
			// 随机种子
			rand.Seed(time.Now().UnixNano())
			// Fisher-Yates 算法随机打乱数组
			for i := len(newUidPool) - 1; i > 0; i-- {
				j := rand.Intn(i + 1)
				newUidPool[i], newUidPool[j] = newUidPool[j], newUidPool[i]
			}

			if _, err := redis.GetRDB().RPush(ctx, cacheKey, newUidPool).Result(); err != nil {
				l.Errorf("Failed to push uid pool item to cache, cacheKey: %s, item: %+v, err: %s", cacheKey, newUidPool, err)
				return -1, pkgError.RedisInternalError
			}

			newUid = curMaxUid + 1
		} else {
			l.Errorf("Failed to LPop new uid in cache, cacheKey: %s, err: %s", cacheKey, err)
			return -1, pkgError.RedisInternalError
		}
	} else {
		newUid, err = strconv.ParseInt(uidVal, 10, 64)
		if err != nil {
			l.Errorf("Incorrect uid fetched from cache, uidVal: %s, err: %s", uidVal, err)
			return -1, pkgError.InternalError
		}
	}

	return newUid, nil
}
