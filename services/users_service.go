package services

import (
	"errors"
	"golang-jwt/models"
	"golang-jwt/repositories"
	"golang-jwt/utils"
)

type UserService interface {
	AttemptLogin(input models.LoginInput) (models.User, error)
	Register(input models.UserInput) (models.User, error)
	GetUserById(userID uint) (models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func (s *userService) GetUserById(userID uint) (models.User, error) {
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) AttemptLogin(input models.LoginInput) (models.User, error) {
	user, err := s.userRepository.FindUserByEmail(input.Email)
	if err != nil {
		return user, err
	}

	comparedPassword := utils.ComparePassword([]byte(input.Password), []byte(user.Password))
	if !comparedPassword {
		return user, errors.New("password is not valid")
	}

	return user, nil
}

func (s *userService) Register(input models.UserInput) (models.User, error) {
	user, _ := s.userRepository.FindUserByEmail(input.Email)
	if user.ID != 0 {
		return user, errors.New("email is exist")
	}

	user = models.User{}
	user.Email = input.Email
	user.FullName = input.Fullname
	user.Password = utils.HashPassword(input.Password)

	user, err := s.userRepository.CreateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func ProvideUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}