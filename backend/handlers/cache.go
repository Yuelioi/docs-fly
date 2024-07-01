package handlers

import (
	"context"
	"runtime"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// 全局 Redis 客户端
var rdb *redis.Client
var ctx = context.Background()
var globalCache sync.Map

func getOs() int {

	switch os := runtime.GOOS; os {
	case "windows":
		return 0
	default:
		return 1
	}
}

func init() {
	// 初始化 Redis 客户端 仅限linux
	if getOs() == 1 {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // Redis 服务地址
			Password: "",               // Redis 密码
			DB:       0,                // Redis 数据库
		})
	}

}

func saveCache(key string, data interface{}) {
	if getOs() == 1 {
		rdb.Set(ctx, key, data, 10*time.Minute)
	} else {
		globalCache.Store(key, data)
	}
}

func getCache(key string) (bool, interface{}) {
	if getOs() == 1 {
		cachedData, err := rdb.Get(ctx, key).Result()

		if err == redis.Nil {
			return false, "没有缓存"
		} else if err != nil {
			return false, err
		} else {
			return true, cachedData
		}
	} else {

		if cachedData, found := globalCache.Load(key); found {
			return true, cachedData
		}
		return false, "没有缓存"

	}

}
