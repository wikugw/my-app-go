package models

type Requirement struct {
	ID            uint   `gorm:"primaryKey"`
	RecruitmentID uint   `gorm:"not null;index"`
	Description   string `gorm:"size:255;not null"`
}
