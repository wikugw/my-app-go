package repositories

import "my-app/database"

func GetRequirementByRecruitmentId(id uint) ([]string, error) {
	db := database.DB
	var reqs []string

	if errReq := db.Table("requirements").
		Where("recruitment_id = ?", id).
		Pluck("description", &reqs).Error; errReq != nil {
		return nil, errReq
	}

	return reqs, nil
}
