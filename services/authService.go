package services

import (
	"quiz-3/helper"
	"quiz-3/model"
	"quiz-3/repository"
)

type AuthService interface {
	Authenticate(username, password string) (*model.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (s *authService) Authenticate(username, password string) (*model.User, error) {
	hashedPassword := helper.HashPassword(password)
	user, err := s.userRepository.GetUserByUsernameAndPassword(username, hashedPassword)

	if err != nil {
		return nil, err
	}
	return user, nil
}
