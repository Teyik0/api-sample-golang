// @/db/database.go
package db

import (
	"fmt"
	"os"

	"github.com/Teyik0/api-sample-golang/entities"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Connecting to database...")
	database_url, database_url_set := os.LookupEnv("DATABASE_URL")
	if database_url_set == false {
		log.Panicf("DATABASE_URL is not set: %s, %s", database_url, database_url_set)
	}
	log.Info("| Connected to database : ", database_url)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  database_url,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entities.User{})

	return nil
}
