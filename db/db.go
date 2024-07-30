package db

import (
	"context"
	"mus/logger"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(options *redis.Options) *redis.Client {
	db := redis.NewClient(options)
	if err := db.Ping(context.Background()).Err(); err != nil {
		logger.LogFatal(err)
	}

	return db
}
