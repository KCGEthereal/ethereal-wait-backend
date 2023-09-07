package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

// NewRedis initiates, checks and returns a redis connection
func NewRedis() *redis.Client {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal("Unable to parse Redis URL")
	}

	client := redis.NewClient(opts)

	if err = client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Unable to ping redis")
	}

	log.Println("Connected to redis")
	return client
}
