package models

import (
	"github.com/google/uuid"
)

type Image struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	User      User      `gorm:"foreignKey:UserID"`
	UserID    uuid.UUID `gorm:"not null"`
	ImagePath string    `gorm:"not null"`
	ImageURL  string    `gorm:"not null"`
}
