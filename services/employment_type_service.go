package services

import (
	"my-app/database"
	"my-app/models"
)

func GetAllEmploymentTypes() ([]models.EmploymentType, error) {
	var employmentTypes []models.EmploymentType
	if err := database.DB.Find(&employmentTypes).Error; err != nil {
		return nil, err
	}
	return employmentTypes, nil
}
