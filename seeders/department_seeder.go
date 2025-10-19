package seeders

import (
	"my-app/database"
	"my-app/models"
)

func SeedDepartments() {
	// Hapus semua data department
	database.DB.Exec("DELETE FROM departments")

	// Insert data baru sekaligus
	dummyDepartments := []models.Department{
		{Name: "HR"},
		{Name: "Finance"},
		{Name: "IT"},
		{Name: "Marketing"},
	}

	database.DB.Create(&dummyDepartments)
}
