package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"test-task-jungle-consulting/configuration"
	"test-task-jungle-consulting/internal/core/dtos"
	"test-task-jungle-consulting/internal/core/models"
	"test-task-jungle-consulting/internal/core/validator"
	"time"
)

type UserServiceUserRepository interface {
	GetUserByName(username string) (models.User, error)
}

type UserService struct {
	UserRepository UserServiceUserRepository
}

func NewUserService(userRepository UserServiceUserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) GetTokenByNameAndPassword(input dtos.LoginInput) (dtos.LoginOutput, error) {
	err := validator.ValidateStruct(input)
	if err != nil {
		return dtos.LoginOutput{}, ValidateError{err}
	}

	user, err := s.UserRepository.GetUserByName(input.Name)
	if err != nil {
		return dtos.LoginOutput{}, RepositoryError{err}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return dtos.LoginOutput{}, BcryptError{err}
	}

	token, err := UserJwtToken(user.ID, user.Username)
	if err != nil {
		return dtos.LoginOutput{}, JWTGenerateError{err}
	}

	return dtos.LoginOutput{
		Token: token,
	}, nil
}

func UserJwtToken(id uuid.UUID, name string) (string, error) {

	now := time.Now().UTC()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  id.String(),
		"name": name,
		"exp":  now.Add(configuration.EnvConfig.JwtExpiresIn).Unix(),
		"iat":  now.Unix(),
		"nbf":  now.Unix(),
	})

	return token.SignedString([]byte(configuration.EnvConfig.JwtSecret))

}
