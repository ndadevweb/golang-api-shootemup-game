package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

const (
	SERVICE_NAME = "back-redis"
	SERVICE_PORT = "6379"
)

var RedisClient *redis.Client

func InitRedis() {
	addr := SERVICE_NAME + ":" + SERVICE_PORT

	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("[KO] Unable to connect to Redis Server : %v", err)
	}

	log.Println("[OK] Redis connected on: ", addr)
}
