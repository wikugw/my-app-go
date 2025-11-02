package services

import (
	"my-app/models"
	"my-app/repositories"
	"my-app/types/handlers"
	"time"
)

func CreateRecruitment(r *handlers.RecruitmentParam) error {
	// Mapping ke models.Recruitment
	recruitment := models.Recruitment{
		DepartmentID:         r.DepartmentID,
		Salary:               r.Salary,
		EmploymentTypeID:     r.EmploymentTypeID,
		ApplicationStartDate: r.ApplicationStartDate,
		ApplicationEndDate:   r.ApplicationEndDate,
		CreatedByID:          r.CreatedByID,
	}

	// Mapping ke []models.Requirement
	var requirements []models.Requirement
	for _, req := range r.Requirements {
		requirements = append(requirements, models.Requirement{
			Description: req,
		})
	}

	return repositories.CreateRecruitment(&recruitment, requirements)
}

func GetActiveRecruitments(date time.Time) ([]handlers.RecruitmentResponse, error) {
	return repositories.GetActiveRecruitments(date)
}
