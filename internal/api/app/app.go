package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"test-task-jungle-consulting/configuration"
	"test-task-jungle-consulting/internal/api/controllers"
	"test-task-jungle-consulting/internal/api/middleware"
	"test-task-jungle-consulting/internal/api/routers"
	"test-task-jungle-consulting/internal/core/services"
	"test-task-jungle-consulting/internal/data/database"
	"test-task-jungle-consulting/internal/data/repositories"
)

type Repositories struct {
	UserRepository  *repositories.UserRepository
	ImageRepository *repositories.ImageRepository
}

type Services struct {
	UserService  *services.UserService
	ImageService *services.ImageService
}

type Controllers struct {
	UserController  *controllers.UserController
	ImageController *controllers.ImageController
}

type Routers struct {
	UserRouter  func(router fiber.Router)
	ImageRouter func(router fiber.Router)
}

func Run(cfg *configuration.EnvConfigModel) {
	db := database.ConnectDB(cfg)

	rp := Repositories{
		UserRepository:  repositories.NewUserRepository(db),
		ImageRepository: repositories.NewImageRepository(db),
	}

	s := Services{
		UserService:  services.NewUserService(rp.UserRepository),
		ImageService: services.NewImageService(rp.ImageRepository),
	}

	c := Controllers{
		UserController:  controllers.NewUserController(s.UserService),
		ImageController: controllers.NewImageController(s.ImageService),
	}

	r := Routers{
		UserRouter:  routers.NewUserRouter(c.UserController),
		ImageRouter: routers.NewImageRouter(c.ImageController),
	}

	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AllowOrigins,
		AllowHeaders: cfg.AllowHeaders,
	}))

	r.UserRouter(app)
	r.ImageRouter(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.APIPort)))
}
