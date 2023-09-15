package database

import (
	"fmt"
	"log"
	"os"

	"github.com/KCGEthereal/ethereal-wait-backend/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGorm initiates, checks and returns a gorm connection
func NewGorm() (client *gorm.DB) {
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	if err = client.AutoMigrate(&models.Form{}); err != nil {
		log.Fatal("Unable to migrate tables")
	}

	return
}
