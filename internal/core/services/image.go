package services

import (
	"github.com/google/uuid"
	"test-task-jungle-consulting/internal/core/dtos"
)

type ImageServiceImageRepository interface {
	InsertImage(up dtos.UploadPicture) error
	GetAllImagesByUserId(userId uuid.UUID) ([]dtos.GetImage, error)
	GetImagePathByUrl(url string) (dtos.AccessImage, error)
}

type ImageService struct {
	ImageRepository ImageServiceImageRepository
}

func NewImageService(imageRepository ImageServiceImageRepository) *ImageService {
	return &ImageService{ImageRepository: imageRepository}
}

func (s *ImageService) UploadPicture(up dtos.UploadPicture) error {
	err := s.ImageRepository.InsertImage(up)
	if err != nil {
		return RepositoryError{err}
	}
	return nil
}

func (s *ImageService) GetAllImages(userId uuid.UUID) ([]dtos.GetImage, error) {
	images, err := s.ImageRepository.GetAllImagesByUserId(userId)
	if err != nil {
		return images, RepositoryError{err}
	}
	return images, nil
}

func (s *ImageService) GetImage(url string) (dtos.AccessImage, error) {
	path, err := s.ImageRepository.GetImagePathByUrl(url)
	if err != nil {
		return dtos.AccessImage{}, RepositoryError{err}
	}
	return path, nil
}
