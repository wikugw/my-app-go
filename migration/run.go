package migration

import (
	"log"
	"my-app/database"
	"my-app/models"
	"my-app/seeders"
)

func RunAll() {

	if err := database.DB.AutoMigrate(
		&models.Employee{},
		&models.Department{},
	); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	seeders.SeedDepartments()
	log.Println("🚀 Migration successful!")
}
