package kv

import (
	"github.com/go-redis/redis"
	"github.com/lethe/config"
)

var RedisClient config.RedisLock

func init() {
	RedisClient = config.RedisLock{
		Client: redis.NewClient(&redis.Options{
			Addr: 		"localhost:6379",
			Password: 	"",
			DB: 		0,
			PoolSize: 	100,
		}),
	}
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}
