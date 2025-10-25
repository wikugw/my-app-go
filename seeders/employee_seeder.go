package seeders

import (
	"fmt"
	"my-app/database"
	"my-app/models"
	"time"
)

func SeedEmployee() error {
	// Pakai global database.DB langsung
	db := database.DB

	var itDept models.Department
	if err := db.Where("name = ?", "IT").First(&itDept).Error; err != nil {
		fmt.Println("Department not found:", err)
		return err
	}

	var fullTime models.EmploymentType
	if err := db.Where("name = ?", "Full Time").First(&fullTime).Error; err != nil { // 🩹 perbaiki variable
		fmt.Println("employment type not found:", err)
		return err
	}

	var count int64
	db.Model(&models.Employee{}).Where("email = ?", "wikugalindrawardhana15@gmail.com").Count(&count)
	if count > 0 {
		return nil
	}

	employee := models.Employee{
		FullName:         "Wiku Galindra Wardhana",
		Email:            "wikugalindrawardhana15@gmail.com",
		DepartmentID:     &itDept.ID,
		HireDate:         time.Date(2022, time.January, 10, 0, 0, 0, 0, time.UTC),
		EmploymentTypeID: &fullTime.ID,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	return db.Create(&employee).Error
}
