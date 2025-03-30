package captcha

import (
	"codeReleaseSystem/pkg/logger"
	"codeReleaseSystem/pkg/redis"
	"time"

	redisv8 "github.com/go-redis/redis/v8"
)

type Store interface {
	Set(key string, value string, expire int) error
	Get(key string) (string, error)
	Delete(key string) error
}

type RedisStore struct{}

func NewRedisStore() *RedisStore {
	return &RedisStore{}
}

func (r *RedisStore) Set(key string, value string, expire int) error {
	if !redis.Redis.Set("captcha:"+key, value, time.Duration(expire)*time.Second) {
		logger.ErrorString("Captcha", "RedisSet", "failed to set captcha")
		return redisv8.Nil
	}
	return nil
}

func (r *RedisStore) Get(key string) (string, error) {
	result := redis.Redis.Get("captcha:" + key)
	if result == "" {
		return "", redisv8.Nil
	}
	return result, nil
}

func (r *RedisStore) Delete(key string) error {
	if !redis.Redis.Del("captcha:" + key) {
		return redisv8.Nil
	}
	return nil
}
