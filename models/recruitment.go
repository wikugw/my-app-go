package models

import "time"

type Recruitment struct {
	ID                   uint       `gorm:"primaryKey"`
	DepartmentID         *uint      `gorm:"column:department_id"`
	Salary               float64    `gorm:"not null"`
	EmploymentTypeID     *uint      `gorm:"column:employment_type_id"`
	ApplicationStartDate *time.Time `gorm:"type:date"`
	ApplicationEndDate   *time.Time `gorm:"type:date"`
	CreatedByID          *uint      `gorm:"column:created_by_id"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
