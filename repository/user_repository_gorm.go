package repository

import (
	"go-rest-api/models"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{DB: db}
}

func (g GormUserRepository) CreateUser(user *models.User) error {
	return g.DB.Create(user).Error
}

func (g GormUserRepository) UpdateUser(user *models.User) error {
	return g.DB.Model(user).Updates(user).Error
}

func (g GormUserRepository) DeleteUser(user *models.User) error {
	return g.DB.Delete(user).Error
}

func (g GormUserRepository) GetUserById(userId int) (*models.User, error) {
	var user models.User
	if err := g.DB.First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g GormUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := g.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g GormUserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := g.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g GormUserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := g.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}
