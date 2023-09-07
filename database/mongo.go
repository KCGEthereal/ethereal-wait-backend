package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func NewMongo() (client *mongo.Client) {
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	var uri string
	if os.Getenv("APP_STAGE") == "production" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin&replicaSet=rs0", user, pass, host, port, name)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", user, pass, host, port, name)
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Unable to connect to mongodb")
	}

	log.Println("Connected to mongodb:", name)

	return client
}
