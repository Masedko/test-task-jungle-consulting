package routers

import (
	"github.com/gofiber/fiber/v2"
	"test-task-jungle-consulting/internal/api/controllers"
	"test-task-jungle-consulting/internal/api/middleware"
)

func NewImageRouter(c *controllers.ImageController) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post("/upload-picture", middleware.Protected(), c.UploadPicture)
		router.Get("/images", middleware.Protected(), c.Images)
		router.Get("/image/:imageUrl", c.Image)
	}
}
