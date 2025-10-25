package seeders

import (
	"my-app/database"
	"my-app/models"
)

func SeedDepartments() {
	db := database.DB

	// Hapus semua data dulu
	db.Exec("DELETE FROM departments")

	departments := []models.Department{
		{ID: 1, Name: "HR"},
		{ID: 2, Name: "Finance"},
		{ID: 3, Name: "IT"},
		{ID: 4, Name: "Marketing"},
	}

	db.Create(&departments)
}
