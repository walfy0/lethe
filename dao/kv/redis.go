package kv

import "github.com/go-redis/redis"

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: 		"localhost:6379",
		Password: 	"",
		DB: 		0,
		PoolSize: 	100,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}
