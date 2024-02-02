package config

import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB{
	var err error
	dsn := os.Getenv("DATABASE_URL")

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	return db
}