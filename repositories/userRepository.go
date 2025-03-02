package repositories

import (
	"errors"
	"github/tiagoduarte/golang-api/database"
	"github/tiagoduarte/golang-api/models"
	"log"
)

func GetUsers(offset int, recordPerPage int) ([]models.User, error) {
	var users []models.User
	if err := database.DB.Offset(offset).Limit(recordPerPage).Order("id ASC").Find(&users).Error; err != nil {
		return nil, errors.New("notfound")
	}

	return users, nil
}

func GetUserByID(userID int) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return user, errors.New("notfound")
	}

	return user, nil
}

func UpdateUser(user *models.User) (*models.User, error) {

	if err := database.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(user *models.User) error {
log.Println("user repositorie:", user)
	if err := database.DB.Delete(user).Error; err != nil {
		return err
	}

	return nil
}
