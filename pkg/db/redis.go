package db

import "github.com/go-redis/redis/v8"

const PoolSize = 20

func NewRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
		PoolSize: PoolSize,
	})
}
