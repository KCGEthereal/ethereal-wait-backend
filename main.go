package main

import (
	"github.com/esportsclub/entity-service-golang/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if stage := os.Getenv("APP_STAGE"); stage != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Unable to load ENV file")
		}
	}
}

func main() {
	mongo := database.NewMongo()
	redis := database.NewRedis()

	err := NewApp(mongo, redis).
		SetupRoutes().
		Listen()
	if err != nil {
		log.Fatal("Unable to listen to port :3000")
	}
}
