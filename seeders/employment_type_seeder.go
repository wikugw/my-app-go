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
		{ID: 1, Name: "Full Time"},
		{ID: 2, Name: "Contract"},
		{ID: 3, Name: "Part Time"},
	}

	database.DB.Create(&dummyEmploymentTypes)
}
