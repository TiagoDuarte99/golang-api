package repositories

import (
	"errors"
	"github/tiagoduarte/golang-api/database"
	"github/tiagoduarte/golang-api/models"
)

func Signup(user *models.User) (error) {
	if err := database.DB.Create(&user).Error; err != nil {
		return errors.New("error creating user")
	}
	return nil
}

func GetUserByEmailForLogin(userEmail string) (models.User, error) {
	var existingUser models.User
	if err := database.DB.Where("email = ?", userEmail).First(&existingUser).Error; err != nil {

		return existingUser, errors.New("user not found")
	}

	return existingUser, nil
}

func CheckIfEmailExists(userEmail string) error {
	var existingUser models.User

	if err := database.DB.Where("email = ?", userEmail).First(&existingUser).Error; err == nil {

		return errors.New("email already registered")
	}

	return nil
}