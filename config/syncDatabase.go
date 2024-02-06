package config

import (
	"log"
	"github.com/iamtaufik/coursehub/models"
)

// SyncDatabase is a function to sync the database
func SyncDatabase() {
	err := DB.AutoMigrate(
		&models.User{}, 
		&models.Course{}, 
		// &models.Category{},
		// &models.Course{},
		// &models.Chapter{},
		// &models.Module{},
	)
	if err != nil {
		log.Fatal("Error migrating the database. Error: ", err)
	}
}