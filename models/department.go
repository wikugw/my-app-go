package models

import "time"

type Department struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
