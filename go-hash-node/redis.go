package main

import (
	"github.com/go-redis/redis"
)

var (
	redisClient *redis.Client
)

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis.prv:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
