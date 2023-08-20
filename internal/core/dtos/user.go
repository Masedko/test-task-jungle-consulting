package dtos

type LoginInput struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}

type LoginOutput struct {
	Token string
}
