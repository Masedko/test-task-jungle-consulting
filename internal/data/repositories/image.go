package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"test-task-jungle-consulting/internal/core/dtos"
	"test-task-jungle-consulting/internal/core/models"
)

type ImageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{db: db}
}

func (r *ImageRepository) InsertImage(up dtos.UploadPicture) error {
	record := r.db.
		Table("images").
		Create(&up)
	return record.Error
}

func (r *ImageRepository) GetAllImagesByUserId(userId uuid.UUID) ([]dtos.GetImage, error) {
	var images []dtos.GetImage
	record := r.db.
		Model(&models.Image{}).
		Where("user_id = ?", userId).
		Find(&images)
	return images, record.Error
}

func (r *ImageRepository) GetImagePathByUrl(url string) (dtos.AccessImage, error) {
	var accessImage dtos.AccessImage
	record := r.db.
		Model(&models.Image{}).
		Where("image_url = ?", url).
		First(&accessImage)
	return accessImage, record.Error
}
