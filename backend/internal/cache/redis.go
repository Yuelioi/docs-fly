package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct{}

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (rc *Redis) Put(ctx context.Context, key, value string, expire time.Duration) (err error) {
	err = rdb.Set(ctx, key, value, expire).Err()
	return err
}
func (rc *Redis) Get(ctx context.Context, key string) (result string, err error) {
	result, err = rdb.Get(ctx, "key1").Result()
	return
}
