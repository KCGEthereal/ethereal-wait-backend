package main

import (
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
	//
}
