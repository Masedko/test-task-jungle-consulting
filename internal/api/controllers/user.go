package controllers

type UserService interface {
}

type UserController struct {
	UserService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{UserService: userService}
}
