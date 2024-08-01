package repository

import "go-rest-api/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
	GetUserById(userId int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
}
