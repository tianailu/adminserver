package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	pkgError "github.com/tianailu/adminserver/pkg/errors"
	"github.com/tianailu/adminserver/pkg/utility/stringx"
	"sync"
	"time"
)

//RedLock The RedLock is a redis red lock.
type RedLock struct {
	clients        []*redis.Client // Redis客户端
	successClients []*redis.Client // 加锁成功的客户端
	unlockScript   *script         // 解锁脚本
	resource       string          // 锁定的资源
	randomValue    string          // 随机值
	watchDog       chan struct{}   // 看门狗
	ttl            time.Duration   // 过期时间
}

//NewRedLock return a new RedLock instance.
func NewRedLock(clients []*redis.Client, resource string, ttl time.Duration) *RedLock {
	return &RedLock{
		clients:      clients,
		unlockScript: unlockScript,
		resource:     resource,
		watchDog:     make(chan struct{}),
		ttl:          ttl,
	}
}

func (rl *RedLock) TryLock(ctx context.Context) error {
	randomValue := stringx.RandN(randomLen)
	var wg sync.WaitGroup
	wg.Add(len(rl.clients))

	// 成功获得锁的Redis实例的客户端
	successClients := make(chan *redis.Client, len(rl.clients))
	for _, client := range rl.clients {
		go func(client *redis.Client) {
			defer wg.Done()
			success, err := client.SetNX(ctx, rl.resource, randomValue, rl.ttl).Result()
			if err != nil {
				return
			}
			// 加锁失败
			if !success {
				return
			}
			// 加锁成功，启动看门狗
			go rl.startWatchDog()
			successClients <- client
		}(client)
	}

	// 等待所有获取锁操作完成
	wg.Wait()
	close(successClients)

	// 如果成功加锁得客户端少于客户端数量的一半+1，表示加锁失败
	if len(successClients) < len(rl.clients)/2+1 {
		// 就算加锁失败，也要把已经获得的锁给释放掉
		for client := range successClients {
			go func(client *redis.Client) {
				ctx, cancel := context.WithTimeout(context.Background(), rl.ttl)
				rl.unlockScript.Run(ctx, client, []string{rl.resource}, randomValue)
				cancel()
			}(client)
		}
		return pkgError.RedisTryLockFailed
	}

	// 加锁成功，启动看门狗
	rl.randomValue = randomValue
	rl.successClients = nil
	for successClient := range successClients {
		rl.successClients = append(rl.successClients, successClient)
	}

	return nil
}

func (rl *RedLock) Unlock(ctx context.Context) error {
	for _, client := range rl.successClients {
		go func(client *redis.Client) {
			rl.unlockScript.Run(ctx, client, []string{rl.resource}, rl.randomValue)
		}(client)
	}

	// 关闭看门狗
	close(rl.watchDog)

	return nil
}

func (rl *RedLock) startWatchDog() {
	rl.watchDog = make(chan struct{})
	resetTTLInterval := rl.ttl / 3
	ticker := time.NewTicker(resetTTLInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 延长锁的过期时间
			for _, client := range rl.successClients {
				go func(client *redis.Client) {
					ctx, cancel := context.WithTimeout(context.Background(), rl.ttl-resetTTLInterval)
					client.Expire(ctx, rl.resource, rl.ttl)
					cancel()
				}(client)
			}
		case <-rl.watchDog:
			// 已经解锁
			return
		}
	}
}
