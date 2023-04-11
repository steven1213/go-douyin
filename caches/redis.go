package caches

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

// Init 初始化redis
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     30,
		PoolTimeout:  30 * time.Second,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
