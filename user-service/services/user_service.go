package services

import (
	"errors"
	"user-service/models"
	"user-service/storage"
	"user-service/utils"
)

type UserService struct {
	Storage storage.StorageInterface
}

func NewUserService(storage storage.StorageInterface) UserServiceInterface {
	return &UserService{Storage: storage}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Storage.GetAllUsers()
}

func (s *UserService) CreateUser(user models.User) error {
	if !utils.ValidateEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if user.Name == "" {
		return errors.New("name is required")
	}

	existingUser, err := s.Storage.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("email already in use")
	}

	return s.Storage.CreateUser(user)
}

func (s *UserService) GetUserById(id string) (*models.User, error) {
	user, err := s.Storage.GetUserById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *UserService) UpdateUser(user models.User) error {
	return s.Storage.UpdateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.Storage.DeleteUser(id)
}
