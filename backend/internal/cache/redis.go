package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	rdb *redis.Client
}

var RC *RedisCache

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	RC = &RedisCache{rdb: rdb}
}

func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) (err error) {
	return rc.rdb.Set(ctx, key, value, expire).Err()

}
func (rc *RedisCache) Get(ctx context.Context, key string) (result string, err error) {
	result, err = rc.rdb.Get(ctx, "key1").Result()
	return
}
