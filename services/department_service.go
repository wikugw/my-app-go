package services

import (
	"errors"
	"my-app/database"
	"my-app/models"
)

var (
	ErrDepartmentExists   = errors.New("department already exists")
	ErrDepartmentNotFound = errors.New("department not found")
)

func GetAllDepartments() ([]models.Department, error) {
	var depts []models.Department
	if err := database.DB.Find(&depts).Error; err != nil {
		return nil, err
	}
	return depts, nil
}
