package redis

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"github.com/tianailu/adminserver/config"
)

const Nil = redis.Nil

// redis 全局实例
var (
	rdb *redis.Client
)

func GetRDB() *redis.Client {
	return rdb
}

// InitRedis 初始化redis
func InitRedis(conf config.Redis) {
	redisUrl := fmt.Sprintf("redis://%s:%s@%s:%d/%d", conf.Username, conf.Password, conf.Ip, conf.Port, conf.DB)

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

	log.Infof("Redis successfully connected")
}
