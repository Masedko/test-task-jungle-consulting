package services

type ImageServiceImageRepository interface {
}

type ImageService struct {
	ImageRepository ImageServiceImageRepository
}

func NewImageService(imageRepository ImageServiceImageRepository) *ImageService {
	return &ImageService{ImageRepository: imageRepository}
}
