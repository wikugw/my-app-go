package models

import "time"

type Employee struct {
	ID               uint      `gorm:"primaryKey"`
	FullName         string    `gorm:"size:150;not null"`
	Email            string    `gorm:"uniqueIndex;not null"`
	DepartmentID     *uint     `gorm:"column:department_id"`
	HireDate         time.Time `gorm:"not null"`
	EmploymentTypeID *uint     `gorm:"column:employment_type_id"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
