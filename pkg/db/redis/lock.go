package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	pkgError "github.com/tianailu/adminserver/pkg/errors"
	"github.com/tianailu/adminserver/pkg/utility/stringx"
	"time"
)

//Lock The Lock is a redis distributed lock.
type Lock struct {
	client          *redis.Client // redis客户端
	unlockScript    *script       // 解锁脚本
	resource        string        // 锁定的资源
	randomValue     string        // 随机值
	watchDog        chan struct{} // 看门狗
	ttl             time.Duration // 过期时间
	tryLockInterval time.Duration // 重新获取锁间隔
}

//NewLock return a new Lock instance.
func NewLock(client *redis.Client, resource string, ttl, tryLockInterval time.Duration) *Lock {
	return &Lock{
		client:          client,
		unlockScript:    unlockScript,
		resource:        resource,
		randomValue:     stringx.RandN(randomLen),
		watchDog:        make(chan struct{}),
		ttl:             ttl,
		tryLockInterval: tryLockInterval,
	}
}

func (rl *Lock) Lock(ctx context.Context) error {
	err := rl.TryLock(ctx)
	if err != nil {
		return err
	}
	if !errors.Is(err, pkgError.RedisTryLockFailed) {
		return err
	}

	// 加锁失败，不但尝试
	ticker := time.NewTicker(rl.tryLockInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			// 超时
			return pkgError.RedisTryLockTimeout
		case <-ticker.C:
			err := rl.TryLock(ctx)
			if err != nil {
				return err
			}
			if !errors.Is(err, pkgError.RedisTryLockFailed) {
				return err
			}
		}
	}

}

func (rl *Lock) TryLock(ctx context.Context) error {
	ok, err := rl.client.SetNX(ctx, rl.resource, rl.randomValue, rl.ttl).Result()
	if err != nil {
		return err
	}

	if !ok {
		// 加锁失败
		return pkgError.RedisTryLockFailed
	}

	// 加锁成功，启动看门狗
	go rl.startWatchDog()
	return nil
}

func (rl *Lock) Unlock(ctx context.Context) error {
	err := rl.unlockScript.Run(ctx, rl.client, []string{rl.resource}, rl.randomValue).Err()

	// 关闭看门狗
	close(rl.watchDog)
	return err
}

func (rl *Lock) startWatchDog() {
	resetTTLInterval := rl.ttl / 3
	ticker := time.NewTicker(resetTTLInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 延长锁的过期时间
			ctx, cancel := context.WithTimeout(context.Background(), rl.ttl-resetTTLInterval)
			ok, err := rl.client.Expire(ctx, rl.resource, rl.ttl).Result()
			cancel()
			// 异常或锁已经不存在则不再续期
			if err != nil || !ok {
				return
			}
		case <-rl.watchDog:
			// 已经解锁
			return
		}
	}
}
