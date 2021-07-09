package config

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/xid"
)

type Lock struct{
	isFirst 	bool
	needUnlock 	bool
	key 		string
	cache 		*redis.Client
}

func (l *Lock) Lock(key string, expire time.Duration) (bool, error){
	if l.isFirst {
		return false, fmt.Errorf("repeat lock")
	}
	l.isFirst = true
	l.key = key
	uuid := xid.New().String()
	locked, err := l.cache.SetNX(l.key, uuid, expire).Result()
	if locked {
		l.needUnlock = true
	}
	return locked, err
}

func (l *Lock) Unlock() error {
	if !l.needUnlock {
		return nil
	}
	return l.cache.Del(l.key).Err()
}