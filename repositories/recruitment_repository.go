package repositories

import (
	"my-app/database"
	"my-app/models"

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
