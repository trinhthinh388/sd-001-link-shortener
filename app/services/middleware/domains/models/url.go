package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	UserID    string         `gorm:"not null;type:text"`
	ShortURL  string         `gorm:"primaryKey;type:varchar;size:20"`
	LongURL   string         `gorm:"type:text"`
	CreatedAt time.Time      `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
