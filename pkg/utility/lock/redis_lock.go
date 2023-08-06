package lock

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/tianailu/adminserver/pkg/db/redis"
	"time"
)

func RetryLock(ctx context.Context, key string, expiredSeconds, retryTimes int, retryInternal time.Duration) bool {
	lock := redis.NewRedisLock(redis.GetRDB(), key)
	lock.SetExpire(expiredSeconds)

	for i := 1; i <= retryTimes; i++ {
		log.Infof("The %d times to try to acquire redis distributed lock, key: %s, expiredSeconds: %d", i, key, expiredSeconds)

		val, err := lock.AcquireCtx(ctx)
		if err != nil {
			log.Errorf("Failed to acquire redis distributed lock, time: %d, key: %s, error: %s", i, key, err.Error())
			return false
		}

		if val {
			log.Infof("Try to acquire redis distributed lock successfully, time: %d, key: %s, expiredSeconds: %d", i, key, expiredSeconds)
			return true
		}

		log.Infof("The lock is currently in use by someone else, wait %d micro second and try again, times: %d, key: %s", retryInternal.Microseconds(), i, key)
		time.Sleep(retryInternal)
	}

	return false
}

func ReleaseLock(ctx context.Context, key string) bool {
	lock := redis.NewRedisLock(redis.GetRDB(), key)

	val, err := lock.ReleaseCtx(ctx)
	if err != nil {
		log.Errorf("Failed to release redis distributed lock, key: %s, error: %s", key, err.Error())
		return false
	}

	return val
}
