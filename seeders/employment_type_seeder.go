package seeders

import (
	"my-app/database"
	"my-app/models"
)

func SeedEmploymentType() {
	// Hapus semua data department
	database.DB.Exec("DELETE FROM employment_types")

	// Insert data baru sekaligus
	dummyEmploymentTypes := []models.EmploymentType{
		{Name: "Full Time"},
		{Name: "Contract"},
		{Name: "Part Time"},
	}

	database.DB.Create(&dummyEmploymentTypes)
}
