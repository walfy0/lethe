package config

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisLock struct{
	*redis.Client
}

func (l *RedisLock) GetLock(key,value string, expire time.Duration) {
	for {
		result := l.SetNX(key, value, expire) 
		if result.Val() {
			break
		}else{
			time.Sleep(time.Second*10)
		}
	}
}

func (l *RedisLock) ReleaseLock(key, value string) {
	l.Eval(`
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KETS[1])
		else
			return 0
		end
	`, []string{key}, value)
}