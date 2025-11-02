package repositories

import (
	"my-app/database"
	"my-app/models"
	"my-app/types/handlers"
	"time"

	"gorm.io/gorm"
)

func CreateRecruitment(recruitment *models.Recruitment, requirements []models.Requirement) error {
	db := database.DB

	return db.Transaction(func(tx *gorm.DB) error {
		// 1️⃣ Simpan recruitment utama
		if err := tx.Create(recruitment).Error; err != nil {
			return err
		}

		// 2️⃣ Isi foreign key dari recruitment.ID
		for i := range requirements {
			requirements[i].RecruitmentID = recruitment.ID
		}

		// 3️⃣ Simpan semua requirement
		if len(requirements) > 0 {
			if err := tx.Create(&requirements).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func GetActiveRecruitments(date time.Time) ([]handlers.RecruitmentResponse, error) {
	db := database.DB
	var recruitments []handlers.RecruitmentResponse

	err := db.Table("recruitments").
		Select(`
			recruitments.id,
			departments.name AS position,
			recruitments.salary,
			employment_types.name AS employment_type,
			recruitments.application_start_date,
			recruitments.application_end_date,
			employees.email AS created_by,
			recruitments.created_at
		`).
		Joins("LEFT JOIN departments ON departments.id = recruitments.department_id").
		Joins("LEFT JOIN employment_types ON employment_types.id = recruitments.employment_type_id").
		Joins("LEFT JOIN employees ON employees.id = recruitments.created_by_id").
		Where("recruitments.application_start_date <= ? AND recruitments.application_end_date >= ?", date, date).
		Find(&recruitments).Error

	if err != nil {
		return nil, err
	}

	for i := range recruitments {
		// Ambil requirements
		var reqs []string
		if errReq := db.Table("requirements").
			Where("recruitment_id = ?", recruitments[i].ID).
			Pluck("description", &reqs).Error; errReq != nil {
			return nil, errReq
		}
		recruitments[i].Requirements = reqs

		// Format tanggal ke string slice
		recruitments[i].ApplicationDates = []string{
			recruitments[i].ApplicationStartDate.Format("2006-01-02"),
			recruitments[i].ApplicationEndDate.Format("2006-01-02"),
		}
	}

	return recruitments, nil
}
