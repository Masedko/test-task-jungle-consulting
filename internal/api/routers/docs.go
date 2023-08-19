package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func NewDocsRouter() func(router fiber.Router) {
	return func(router fiber.Router) {
		// Swagger docs
		router.Get("/docs/*", swagger.HandlerDefault)
		// Redirect to docs
		router.Get("/", func(ctx *fiber.Ctx) error {
			return ctx.Redirect("/docs/")
		})
	}
}
