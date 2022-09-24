package repositories

import (
	"golang-jwt/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(userID uint) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) GetUserByID(userID uint) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func ProvideUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}