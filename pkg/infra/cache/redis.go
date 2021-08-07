package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	ctx context.Context
	rdb *redis.Client
}

func logRedisError(text string, err error) {
	log.Println("[REDIS CLIENT ERROR]:", text)
	log.Println("[REDIS CLIENT ERROR]:", err)
}

func NewRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password2021", // no password set
		DB:       0,              // use default DB
	})

	return &RedisClient{
		ctx: context.Background(),
		rdb: rdb,
	}
}

func (rc *RedisClient) Set(key, value string) error {
	err := rc.rdb.Set(rc.ctx, "key", "value", 0).Err()
	if err != nil {
		logRedisError("Error on Set", err)
		return err
	}

	return nil
}

func (rc *RedisClient) Get(key string) (string, error) {
	val, err := rc.rdb.Get(rc.ctx, "key").Result()
	if err != nil {
		logRedisError("Error on Get", err)
		return "", err
	}

	return val, nil
}

func (rc *RedisClient) Del(key string) (int64, error) {
	val, err := rc.rdb.Del(rc.ctx, "key").Result()
	if err != nil {
		logRedisError("Error on Get", err)
		return val, err
	}

	return val, nil
}
