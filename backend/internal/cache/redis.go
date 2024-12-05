package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Rctx context.Context

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	Rctx = context.Background()

}

// func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) (err error) {
// 	return rc.rdb.Set(ctx, key, value, expire).Err()

// }
// func (rc *RedisCache) Get(ctx context.Context, key string) (result string, err error) {
// 	result, err = rc.rdb.Get(ctx, "key1").Result()
// 	return
// }
