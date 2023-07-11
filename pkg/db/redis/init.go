package redis

import (
    "context"
    "fmt"
    "github.com/labstack/gommon/log"
    "github.com/redis/go-redis/v9"
)

// redis全局实例
var (
    rdb *redis.Client
)

func GetRDB() *redis.Client {
    return rdb
}

// InitRedis 初始化redis
func InitRedis(username, password, host, port, db string) {
    redisUrl := fmt.Sprintf("redis://%s:%s@%s:%s/%s", username, password, host, port, db)
    opt, err := redis.ParseURL(redisUrl)
    if err != nil {
        log.Errorf("生成 Redis 访问URL异常: %s", err.Error())
        return
    }

    rdb = redis.NewClient(opt)

    ping := rdb.Ping(context.TODO())
    if _, err = ping.Result(); err != nil {
        log.Errorf("Redis Ping err: %s", err.Error())
        return
    }

    fmt.Println("Redis successfully connected")
}
