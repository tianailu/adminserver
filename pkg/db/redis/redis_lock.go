package redis

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"github.com/tianailu/adminserver/pkg/utility/stringx"
	"strconv"
	"sync/atomic"
)

const (
	randomLen       = 16
	tolerance       = 500 // milliseconds
	millisPerSecond = 1000
)

var (
	lockScript = NewScript(`if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "OK"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`)
	delScript = NewScript(`if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end`)
)

// A RedisLock is a redis lock.
type RedisLock struct {
	rdb     *redis.Client
	seconds uint32
	key     string
	id      string
}

type (
	Script = redis.Script
)

// NewScript returns a new Script instance.
func NewScript(script string) *Script {
	return redis.NewScript(script)
}

func NewRedisLock(rdb *redis.Client, key string) *RedisLock {
	return &RedisLock{
		rdb: rdb,
		key: key,
		id:  stringx.RandN(randomLen),
	}
}

// Acquire acquires the lock.
func (rl *RedisLock) Acquire() (bool, error) {
	return rl.AcquireCtx(context.Background())
}

// AcquireCtx acquires the lock with the given ctx.
func (rl *RedisLock) AcquireCtx(ctx context.Context) (bool, error) {
	seconds := atomic.LoadUint32(&rl.seconds)

	val, err := lockScript.Run(ctx, rl.rdb, []string{rl.key}, []string{
		rl.id, strconv.Itoa(int(seconds)*millisPerSecond + tolerance),
	}).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		log.Errorf("Error on acquiring lock for %s, %s", rl.key, err.Error())
		return false, err
	} else if val == nil {
		return false, nil
	}

	reply, ok := val.(string)
	if ok && reply == "OK" {
		return true, nil
	}

	log.Errorf("Unknown reply when acquiring lock for %s: %v", rl.key, val)
	return false, nil
}

// Release releases the lock.
func (rl *RedisLock) Release() (bool, error) {
	return rl.ReleaseCtx(context.Background())
}

// ReleaseCtx releases the lock with the given ctx.
func (rl *RedisLock) ReleaseCtx(ctx context.Context) (bool, error) {
	val, err := delScript.Run(ctx, rl.rdb, []string{rl.key}, []string{rl.id}).Result()
	if err != nil {
		return false, err
	}

	reply, ok := val.(int64)
	if !ok {
		return false, nil
	}

	log.Infof("Release redis lock, key: %s, val: %s, reply: %d", rl.key, rl.id, reply)

	return reply == 1, nil
}

// SetExpire sets the expiration.
func (rl *RedisLock) SetExpire(seconds int) {
	atomic.StoreUint32(&rl.seconds, uint32(seconds))
}
