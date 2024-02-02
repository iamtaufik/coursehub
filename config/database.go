package config

import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")

	// Connect to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}
}