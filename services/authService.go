package services

import (
	"fmt"
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
	fmt.Println("INi hashed password: ", hashedPassword)
	user, err := s.userRepository.GetUserByUsernameAndPassword(username, hashedPassword)
	fmt.Println("Cek user by username dan password: ", user)

	if err != nil {
		return nil, err
	}
	return user, nil
}
