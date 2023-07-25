package gold

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"github.com/tianailu/adminserver/pkg/db/redis"
	"log"
	"time"
)

type DBSlice[T GoldTradeList | InviteInfo] []T

func (d *DBSlice[T]) GetOnePageWithOrder(page, size int, table string) (int64, error) {
	var num int64
	result := mysql.GetDB().Table(table).Limit(size).Offset(page).Find(&d).Where("is_del = ?", 1).Count(&num)
	if result.Error != nil {
		return 0, result.Error
	}
	return num / int64(size), nil
}

func (d *DBSlice[T]) GetOnePageWithOrderAndWhere(page, size int, table string, where map[string]interface{}) error {
	result := mysql.GetDB().Table(table).Where(where).Limit(size).Offset(page).Find(&d)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *DBSlice[T]) Set(table string) error {
	result := mysql.GetDB().Table(table).Create(&d)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type DBValue[T ChatSetting | GoldSetting] struct {
	Data T
}

func (d *DBValue[T]) SetWithCache(ctx context.Context, table string, param, value string) error {
	// 删除redis
	if err := redis.GetRDB().Del(ctx, getRedisKey(table)).Err(); err != nil {
		log.Println(fmt.Sprintf("[SetWithCache] del redis key fail,err:%v", err))
		return err
	}
	if err := mysql.GetDB().Table(table).Update(param, value).Error; err != nil {
		log.Println(fmt.Sprintf("[SetWithCache] Update MySQL key fail,err:%v", err))
		return err
	}
	// 延迟双删
	go func() {
		time.Sleep(10 * time.Second)
		for i := 0; i < 3; i++ {
			if err := redis.GetRDB().Del(ctx, getRedisKey(table)).Err(); err == nil {
				return
			}
		}

	}()
	return nil
}

func (d *DBValue[T]) GetSetting(ctx context.Context, table string) error {
	rdbResult, err := redis.GetRDB().Get(ctx, getRedisKey(table)).Result()
	if err != nil {
		if err := mysql.GetDB().Table(table).First(&d.Data).Error; err != nil {
			log.Println(fmt.Sprintf("[GetSetting] query MySQL key fail,err:%v", err))
			return err
		}
		rdbValue, err := json.Marshal(d.Data)
		if err != nil {
			log.Println(fmt.Sprintf("[GetSetting] marshal json is fail, err:%v", err))
			return nil
		}
		if err := redis.GetRDB().Set(ctx, getRedisKey(table), rdbValue, -1).Err(); err != nil {
			log.Println(fmt.Sprintf("[GetSetting] set redis cache is fail, err:%v", err))
			return nil
		}
	}
	if err := json.Unmarshal([]byte(rdbResult), d.Data); err != nil {
		log.Println(fmt.Sprintf("[GetSetting] Unmarshal json is fail, err:%v", err))
		return err
	}
	return nil
}
func (d *DBValue[T]) GetSettingWithNameSpace(ctx context.Context, table, namespace string) error {
	rdbResult, err := redis.GetRDB().Get(ctx, getRedisKeyWithNameSpace(table, namespace)).Result()
	if err != nil {
		if err := mysql.GetDB().Table(table).First(&d.Data).Error; err != nil {
			log.Println(fmt.Sprintf("[GetSetting] query MySQL key fail,err:%v", err))
			return err
		}
		rdbValue, err := json.Marshal(d.Data)
		if err != nil {
			log.Println(fmt.Sprintf("[GetSetting] marshal json is fail, err:%v", err))
			return nil
		}
		if err := redis.GetRDB().Set(ctx, getRedisKeyWithNameSpace(table, namespace), rdbValue, -1).Err(); err != nil {
			log.Println(fmt.Sprintf("[GetSetting] set redis cache is fail, err:%v", err))
			return nil
		}
	}
	if err := json.Unmarshal([]byte(rdbResult), d.Data); err != nil {
		log.Println(fmt.Sprintf("[GetSetting] Unmarshal json is fail, err:%v", err))
		return err
	}
	return nil
}

func getRedisKey(table string) string {
	return fmt.Sprintf("admin_%v", table)
}

func getRedisKeyWithNameSpace(table, namespace string) string {
	return fmt.Sprintf("admin_%v_%v", table, namespace)
}
