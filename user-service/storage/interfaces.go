package storage

import "user-service/models"

type StorageInterface interface {
	CreateUser(user models.User) error
	GetUserById(id string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id string) error
	GetUserByEmail(email string) (*models.User, error)
}
