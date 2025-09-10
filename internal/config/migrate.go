package config

import (
	"log"

	"github.com/qullDev/BookStore-API/internal/models"
)

// MigrateDB runs auto migration for all models
func MigrateDB() {

	err := DB.AutoMigrate(
		&models.Author{},
		&models.Category{},
		&models.Book{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Print("Database migration completed successfully")
}
