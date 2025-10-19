package services

import (
	"errors"
	"time"

	"my-app/database"
	"my-app/models"

	employee "my-app/types/services"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	ErrEmployeeExists   = errors.New("email already exists")
	ErrEmployeeNotFound = errors.New("employee not found")
)

func CreateEmployee(req employee.CreateEmployeeRequest) (*models.Employee, error) {
	employee := models.Employee{
		FullName:   req.FullName,
		Email:      req.Email,
		Position:   req.Position,
		Department: req.Department,
		HireDate:   time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := database.DB.Create(&employee).Error; err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			return nil, ErrEmployeeExists
		}
		return nil, err
	}

	return &employee, nil
}

func GetEmployeeByEmail(email string) (*models.Employee, error) {
	var employee models.Employee
	if err := database.DB.Where("email = ?", email).First(&employee).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrEmployeeNotFound
		}
		return nil, err
	}
	return &employee, nil
}
