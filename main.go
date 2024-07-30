package main

import (
	"mus/api"
	"mus/db"
	"mus/logger"

	"github.com/go-redis/redis/v8"
)

func main() {
	logger.SetLogLevel(logger.InfoLogLevel)

	db := db.NewRedisClient(&redis.Options{
		Addr:     "mus-redis:6379",
		Password: "",
		DB:       0,
	})

	mus := api.NewMusAPI("0.0.0.0:8000", db)
	logger.LogFatal(mus.Run())
}
