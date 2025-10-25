package repositories

import (
	employee "my-app/types/repositories"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{}
}

func (r *EmployeeRepository) GetByEmail(email string) (*employee.EmployeeDetail, error) {
	var result employee.EmployeeDetail
	err := r.db.Table("employees").
		Select(`
			employees.id, employees.full_name, employees.email,
			departments.name AS department_name,
			employment_types.name AS employment_type
		`).
		Joins("LEFT JOIN departments ON departments.id = employees.department_id").
		Joins("LEFT JOIN employment_types ON employment_types.id = employees.employment_type_id").
		Where("employees.email = ?", email).
		First(&result).Error
	return &result, err
}
