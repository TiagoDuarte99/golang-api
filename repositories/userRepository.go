package repositories

import (
	"errors"
	"github/tiagoduarte/golang-api/database"
	"github/tiagoduarte/golang-api/models"
)

func GetUsers(offset int, recordPerPage int) ([]models.User, error) {
	var users []models.User
	if err := database.DB.Offset(offset).Limit(recordPerPage).Order("id ASC").Find(&users).Error; err != nil {
		return nil, errors.New("users not found")
	}

	return users, nil
}

func GetUserByID(userID int) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

