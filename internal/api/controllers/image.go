package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"path/filepath"
	"test-task-jungle-consulting/internal/core/dtos"
	"test-task-jungle-consulting/internal/core/validator"
)

type ImageService interface {
	UploadPicture(up dtos.UploadPicture) error
	GetAllImages(userId uuid.UUID) ([]dtos.GetImage, error)
	GetImage(url string) (dtos.AccessImage, error)
}

type ImageController struct {
	ImageService ImageService
}

func NewImageController(imageService ImageService) *ImageController {
	return &ImageController{ImageService: imageService}
}

// UploadPicture
//
//	@Summary		Upload picture
//	@Description	Upload picture for user
//	@Tags			image
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			picture	formData	file					true	"Body with picture"
//	@Success		200		{object}	dtos.MessageResponse	"Image information"
//	@Failure		401		{object}	dtos.MessageResponse	"Couldn't validate JWT"
//	@Failure		422		{object}	dtos.MessageResponse	"Couldn't get image from form"
//	@Failure		500		{object}	dtos.MessageResponse	"Couldn't save image"
//	@Router			/upload-picture [post]
func (cr *ImageController) UploadPicture(c *fiber.Ctx) error {
	user := c.Locals("user")
	userId, err := validator.GetUserIdAndCheckJWT(user)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusUnauthorized).JSON(dtos.MessageResponse{
			Message: "Couldn't validate JWT",
		})
	}

	file, err := c.FormFile("picture")
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dtos.MessageResponse{
			Message: "Couldn't get image from form",
		})
	}

	imageNewName, err := uuid.NewUUID()
	imageNewNameString := imageNewName.String()
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.MessageResponse{
			Message: "Couldn't generate uuid",
		})
	}

	filePath := StaticDir + imageNewNameString + filepath.Ext(file.Filename)
	urlPath := imageNewNameString + filepath.Ext(file.Filename)
	err = c.SaveFile(file, filePath)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.MessageResponse{
			Message: "Couldn't save image to directory",
		})
	}

	err = cr.ImageService.UploadPicture(dtos.UploadPicture{
		UserID:    userId,
		ImagePath: filePath,
		ImageURL:  urlPath,
	})
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.MessageResponse{
			Message: "Couldn't save image to database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dtos.MessageResponse{
		Message: fmt.Sprintf("Image saved as %s", imageNewNameString),
	})
}

// Images
//
//	@Summary		Get images
//	@Description	Get images for specific user
//	@Tags			image
//	@Produce		json
//	@Security		JWT
//	@Success		200	{object}	[]dtos.GetImage			"Images"
//	@Failure		422	{object}	dtos.MessageResponse	"Couldn't validate JWT"
//	@Failure		500	{object}	dtos.MessageResponse	"Couldn't get images"
//	@Router			/images [get]
func (cr *ImageController) Images(c *fiber.Ctx) error {
	user := c.Locals("user")
	userId, err := validator.GetUserIdAndCheckJWT(user)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusUnauthorized).JSON(dtos.MessageResponse{
			Message: "Couldn't validate JWT",
		})
	}

	images, err := cr.ImageService.GetAllImages(userId)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.MessageResponse{
			Message: "Couldn't get images",
		})
	}

	return c.Status(fiber.StatusOK).JSON(images)
}

// Image
//
//	@Summary		Get image by URL
//	@Description	Retrieve an image from the server using its URL.
//	@Tags			image
//	@Accept			json
//	@Produce		octet-stream
//	@Param			imageURL	path		string					true	"URL of the image"
//	@Success		200			{file}		octet-stream			"Image content"
//	@Failure		500			{object}	dtos.MessageResponse	"Internal server error"
//	@Router			/image/{imageURL} [get]
func (cr *ImageController) Image(c *fiber.Ctx) error {
	imageURL := c.Params("imageURL")
	accessImage, err := cr.ImageService.GetImage(imageURL)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dtos.MessageResponse{
			Message: "Couldn't get path of image",
		})
	}
	return c.Status(fiber.StatusOK).SendFile(accessImage.ImagePath)
}
