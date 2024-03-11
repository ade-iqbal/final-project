package service

import (
	"fga-final-project/helper"
	"fga-final-project/model"
	"fga-final-project/repository"
	"fga-final-project/util"
)

type UserService interface {
	Register(user model.User) (model.User, error)
	Login(user model.User) (string, error)
	UpdateUser(id uint, user model.User) (model.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	UserRepository repository.UserRepository 
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(user model.User) (model.User, error) {
	err := s.UserRepository.CreateUser(&user)
	if err != nil {
		helper.LoggingError("Create Data Error", err)
	}

	return user, err
}

func (s *userService) Login(user model.User) (string, error) {
	userData, err := s.UserRepository.GetUserByEmail(&user.Email)
	if err != nil {
		helper.LoggingError("User Not Found", err)
		return "", util.ErrInvalidCredentials
	}

	comparePass := helper.ComparePassword([]byte(userData.Password), []byte(user.Password))
	if !comparePass {
		return "", util.ErrInvalidCredentials
	}

	token := helper.GenerateToken(userData.ID, userData.Email)
	return token, nil
}

func (s *userService) UpdateUser(id uint, user model.User) (model.User, error) {
	var (
		userData *model.User
		err error
	)

	userData, err = s.UserRepository.GetUserById(&id)
	if err != nil {
		helper.LoggingError("User Not Found", err)
		return *userData, err
	}

	userData.Email = user.Email
	userData.Username = user.Username
	
	err = s.UserRepository.UpdateUser(userData)
	if err != nil {
		helper.LoggingError("Update Data Error", err)
	}

	return *userData, err
}

func (s *userService) DeleteUser(id uint) error {
	err := s.UserRepository.DeleteUser(&id)
	if err != nil {
		helper.LoggingError("Delete Data Error", err)
	}

	return err
}