package models

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `gorm:"not null;default:uuid_generate_v4();primaryKey"`
	Username     string    `gorm:"not null"`
	PasswordHash string    `gorm:"not null"`
}
