package models

import "time"

type Recruitment struct {
	ID               uint      `gorm:"primaryKey"`
	StartDate        time.Time `gorm:"not null"`
	EndDate          time.Time `gorm:"not null"`
	CreatedBy        uint      `gorm:"not null"` // Foreign key to User/Employee
	EmploymentTypeID uint      `gorm:"not null"`
	PositionID       uint      `gorm:"not null"`
	Requirements     string    `gorm:"type:text"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
