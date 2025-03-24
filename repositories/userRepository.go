package repositories

import (
	"github/tiagoduarte/golang-api/database"
	"github/tiagoduarte/golang-api/dto"
	helper "github/tiagoduarte/golang-api/helpers"
	"github/tiagoduarte/golang-api/models"
	"log"
)

func GetUsers(offset int, recordPerPage int) ([]dto.UserResponse, error) {
	var users []models.User
	if err := database.DB.
		Offset(offset).
		Limit(recordPerPage).
		Order("id ASC").
		Find(&users).Error; err != nil {
		return nil, &helper.CustomError{
			Type: helper.ErrNotFound,
			Message: helper.ErrorResponse{
				Message: "Users not found for the requested parameters.",
			},
		}
	}

	var userResponses []dto.UserResponse
	for _, user := range users {
		// Mapeamento de cada usuário para dto.UserResponse
		userResponses = append(userResponses, dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			UserType:  user.UserType,
			CreatedAt: user.CreatedAt,
		})
	}


	return userResponses, nil
}

func GetUserByID(userID int) (*dto.UserResponse, error) {
	var user models.User
	if err := database.DB.
		First(&user, "id = ?", userID).Error; err != nil {
		return nil, &helper.CustomError{
			Type: helper.ErrNotFound,
			Message: helper.ErrorResponse{
				Message: "User not found for the requested parameters.",
			},
		}
	}

	userResponse := &dto.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		UserType:     user.UserType,
		CreatedAt:    user.CreatedAt,
	}

	return userResponse, nil
}

func GetUserCompleteByID(userID int) (*models.User, error) {
	var user models.User
	if err := database.DB.
		First(&user, "id = ?", userID).Error; err != nil {
		return nil, &helper.CustomError{
			Type: helper.ErrNotFound,
			Message: helper.ErrorResponse{
				Message: "User not found for the requested parameters.",
			},
		}
	}

	return &user, nil
}

func UpdateUser(user *models.User) (*dto.UserResponse, error) {
	if err := database.DB.Model(&models.User{}).Where("id = ?", user.ID).
		Select("name", "email", "password").Updates(user).Error; err != nil {
		return nil, err
	}
	
	userResponse := &dto.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		UserType:     user.UserType,
		CreatedAt:    user.CreatedAt,
	}

	return userResponse, nil
}

func DeleteUser(user *dto.UserResponse) error {
	log.Println("user repositorie:", user)
	if err := database.DB.
		Delete(user).Error; err != nil {
		return err
	}

	return nil
}
