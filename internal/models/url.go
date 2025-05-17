package models

import (
	"time"
	"gorm.io/gorm"
)

type URL struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	OriginalURL string `gorm:"not null"`
	ShortCode   string `gorm:"uniqueIndex;not null"`
	Clicks      int
}
