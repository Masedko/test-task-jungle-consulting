package routers

import (
	"github.com/gofiber/fiber/v2"
	"test-task-jungle-consulting/internal/api/controllers"
)

func NewUserRouter(c *controllers.UserController) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post("/login")
	}
}
