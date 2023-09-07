package services

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	DB    *mongo.Client
	Redis *redis.Client
}
