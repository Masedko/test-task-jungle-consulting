package routers

import (
	"github.com/gofiber/fiber/v2"
	"test-task-jungle-consulting/internal/api/controllers"
)

func NewImageRouter(c *controllers.ImageController) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post("/upload-picture")
		router.Get("/images")
	}
}
