package repositories

import (
	"github/tiagoduarte/golang-api/database"
	helper "github/tiagoduarte/golang-api/helpers"

	"github/tiagoduarte/golang-api/models"
)

func Signup(user *models.User) error {
	if err := database.DB.Create(&user).Error; err != nil {
		return &helper.CustomError{
			Type:    helper.ErrConflict,
			Message: helper.ErrorResponse{Message: "Error creating user in database: " + err.Error()},
		}
	}
	return nil
}

func GetUserByEmailForLogin(userEmail string) (models.User, error) {
	var existingUser models.User
	if err := database.DB.
		Where("email = ?", userEmail).
		First(&existingUser).Error; err != nil {

			return existingUser, &helper.CustomError{
				Type:    helper.ErrNotFound,
				Message: helper.ErrorResponse{Message: "User not found"},
			}
	}

	return existingUser, nil
}

func CheckIfEmailExists(userEmail string) error {
	var existingUser models.User

	if err := database.DB.
		Where("email = ?", userEmail).
		First(&existingUser).Error; err == nil {

			return &helper.CustomError{
				Type:    helper.ErrConflict, 
				Message: helper.ErrorResponse{Message: "Email already registered"},
			}
	}

	return nil
}
