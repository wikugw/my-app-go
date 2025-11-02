package handlers

import (
	"time"
)

type CreateRecruitmentRequest struct {
	Position         string   `json:"position" binding:"required"`
	Salary           float64  `json:"salary" binding:"required"`
	EmploymentType   string   `json:"employmentType" binding:"required"`
	ApplicationDates []string `json:"applicationDates" binding:"required,len=2"`
	Requirements     []string `json:"requirements"`
	CreatedByID      uint     `json:"createdById" binding:"required"`
}

type RecruitmentParam struct {
	DepartmentID         uint       `gorm:"column:department_id"`
	Salary               float64    `gorm:"not null"`
	EmploymentTypeID     uint       `gorm:"column:employment_type_id"`
	ApplicationStartDate *time.Time `gorm:"type:date"`
	ApplicationEndDate   *time.Time `gorm:"type:date"`
	CreatedByID          uint       `gorm:"column:created_by_id"`
	Requirements         []string
}

type RecruitmentResponse struct {
	ID                   uint      `json:"id"`
	Position             string    `json:"position"`
	Salary               float64   `json:"salary"`
	EmploymentType       string    `json:"employmentType"`
	ApplicationStartDate time.Time `json:"-"` // hidden in JSON
	ApplicationEndDate   time.Time `json:"-"` // hidden in JSON
	ApplicationDates     []string  `json:"applicationDates"`
	Requirements         []string  `json:"requirements"`
	CreatedBy            string    `json:"createdBy"`
	CreatedAt            string    `json:"createdAt"`
}
