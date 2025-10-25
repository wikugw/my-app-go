package migration

import (
	"log"
	"my-app/database"
	"my-app/models"
	"my-app/seeders"
)

func RunAll() {

	if err := database.DB.AutoMigrate(
		&models.EmploymentType{},
		&models.Employee{},
		&models.Department{},
		&models.Recruitment{},
	); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	seeders.SeedDepartments()
	seeders.SeedEmploymentType()
	seeders.SeedEmployee()
	log.Println("🚀 Migration successful!")
}
