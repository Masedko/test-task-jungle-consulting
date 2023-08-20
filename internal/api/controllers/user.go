package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"test-task-jungle-consulting/internal/core/dtos"
)

type UserService interface {
	GetTokenByNameAndPassword(input dtos.LoginInput) (dtos.LoginOutput, error)
}

type UserController struct {
	UserService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{UserService: userService}
}

// Login
//
//	@Summary		Login
//	@Description	Login
//	@Tags			user
//	@Produce		json
//	@Security		JWT
//	@Param			payload	body		dtos.LoginInput			true	"Data to login"
//	@Success		200		{object}	dtos.LoginOutput		"Bearer token (login information)"
//	@Failure		400		{object}	dtos.MessageResponse	"Couldn't get token"
//	@Failure		422		{object}	dtos.MessageResponse	"Couldn't parse body (login information)"
//	@Router			/login [post]
func (cr *UserController) Login(c *fiber.Ctx) error {
	var payload dtos.LoginInput

	err := c.BodyParser(&payload)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dtos.MessageResponse{
			Message: "Couldn't parse body (login information)",
		})
	}

	details, err := cr.UserService.GetTokenByNameAndPassword(payload)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(dtos.MessageResponse{
			Message: "Couldn't get token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(details)
}
