package services

import (
	"errors"
	"user-service/models"
	"user-service/storage"
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
	return s.Storage.GetUserById(id)
}

func (s *UserService) UpdateUser(user models.User) error {
	existingUser, err := s.Storage.GetUserById(user.ID)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return errors.New("user not found")
	}

	return s.Storage.UpdateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.Storage.DeleteUser(id)
}
