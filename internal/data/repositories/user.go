package repositories

import (
	"gorm.io/gorm"
	"test-task-jungle-consulting/internal/core/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByName(username string) (models.User, error) {
	var user models.User
	record := r.db.
		Model(&models.User{}).
		Where("username = ?", username).
		First(&user)
	return user, record.Error
}
