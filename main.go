package main

import (
	"github.com/esportsclub/entity-service-golang/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// init sole work is to get all the ENV variables from the system and set
// it up to work. APP_STAGE makes sure that .env file is loaded if it's
// set to development. When set to production, we load ENV values directly
// from the system.
func init() {
	if stage := os.Getenv("APP_STAGE"); stage != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Unable to load ENV file")
		}
	}
}

// main handles the bootstrapping and connecting to database.
// We pass around the mongodb connection to App which then
// hands it off to other structs that require a DB connection
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
