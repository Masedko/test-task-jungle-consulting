package services

type UserServiceUserRepository interface {
}

type UserService struct {
	UserRepository UserServiceUserRepository
}

func NewUserService(userRepository UserServiceUserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}
