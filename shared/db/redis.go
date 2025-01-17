package db

import (
	"context"
	"shared/config"
	"shared/helpers"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	cfg, err := config.Load()
	helpers.PanicIfError(err)

	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})

	_, err = redisClient.Ping(context.Background()).Result()
	helpers.PanicIfError(err)

	return redisClient
}
