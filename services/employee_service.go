package services

import (
	"errors"
	"time"

	"my-app/database"
	"my-app/models"
	"my-app/repositories"

	employeeRepo "my-app/types/repositories"
	employee "my-app/types/services"

	"github.com/lib/pq"
)

var (
	ErrEmployeeExists   = errors.New("email already exists")
	ErrEmployeeNotFound = errors.New("employee not found")
)

type EmployeeService struct {
	repo *repositories.EmployeeRepository
}

var EmployeeServiceInstance *EmployeeService

func NewEmployeeService(repo *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func CreateEmployee(req employee.CreateEmployeeRequest) (*models.Employee, error) {
	employee := models.Employee{
		FullName:         req.FullName,
		Email:            req.Email,
		DepartmentID:     req.DepartmentId,
		EmploymentTypeID: req.EmploymentTypeId,
		HireDate:         time.Now(),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := database.DB.Create(&employee).Error; err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			return nil, ErrEmployeeExists
		}
		return nil, err
	}

	return &employee, nil
}

func (s *EmployeeService) GetEmployeeByEmail(email string) (*employeeRepo.EmployeeDetail, error) {
	return s.repo.GetByEmail(email)
}
