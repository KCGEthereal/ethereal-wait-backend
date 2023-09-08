package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"time"
)

type Redis struct {
	ctx    context.Context
	client *redis.Client
}

// NewRedis initiates, checks and returns a Redis struct with
// all of its receiver function
func NewRedis() (redisStruct *Redis) {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal("Unable to parse Redis URL")
	}

	redisStruct.ctx = context.Background()
	redisStruct.client = redis.NewClient(opts)

	if err = redisStruct.client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Unable to ping redis")
	}

	log.Println("Connected to redis")
	return
}

// Remember takes a key and expiration and callback. When the key exists in
// cache, it returns the cached result.
//
// If key does not exist in the cache, it executes the callback function and
// then saves the key with result of the callback while also setting the
// expiration you provided.
func (r *Redis) Remember(key string, expire time.Duration, callback func() string) (string, error) {
	if exists, _ := r.client.Exists(r.ctx, key).Result(); exists == 1 {
		return r.client.Get(r.ctx, key).Result()
	}

	result := callback()
	_, err := r.client.Set(r.ctx, key, result, expire).Result()
	return result, err
}

// Client returns the redis client created by NewRedis
func (r *Redis) Client() *redis.Client {
	return r.client
}
