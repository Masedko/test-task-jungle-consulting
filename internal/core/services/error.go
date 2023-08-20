package services

import (
	"fmt"
)

type ValidateError struct {
	error
}

func (e ValidateError) Error() string {
	return fmt.Sprintf("Validate error: %v", e.error)
}

type UserNotExistsError struct {
	Username string
}

func (e UserNotExistsError) Error() string {
	return fmt.Sprintf("User %s already exists", e.Username)
}

type RepositoryError struct {
	error
}

func (e RepositoryError) Error() string {
	return fmt.Sprintf("Repository Error: %v", e.error)
}

type BcryptError struct {
	error
}

func (e BcryptError) Error() string {
	return fmt.Sprintf("Bcrypt error: %v", e.error)
}

type UUIDError struct {
	error
}

func (e UUIDError) Error() string {
	return fmt.Sprintf("UUID error: %v", e.error)
}

type JWTGenerateError struct {
	error
}

func (e JWTGenerateError) Error() string {
	return fmt.Sprintf("Generating JWT Token failed: %v", e.error)
}
