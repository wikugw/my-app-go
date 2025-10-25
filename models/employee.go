package models

import "time"

type Employee struct {
	ID               uint      `gorm:"primaryKey"`
	FullName         string    `gorm:"size:150;not null"`
	Email            string    `gorm:"uniqueIndex;not null"`
	Position         string    `gorm:"size:100"`
	Department       string    `gorm:"size:100"`
	HireDate         time.Time `gorm:"not null"`
	EmploymentTypeID uint      `gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
