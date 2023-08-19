package controllers

type ImageService interface {
}

type ImageController struct {
	ImageService ImageService
}

func NewImageController(imageService ImageService) *ImageController {
	return &ImageController{ImageService: imageService}
}
