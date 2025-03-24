package services

import (
	"github/tiagoduarte/golang-api/dto"
	helper "github/tiagoduarte/golang-api/helpers"
	repositories "github/tiagoduarte/golang-api/repositories"

	"github/tiagoduarte/golang-api/models"

	"time"
)

func Signup(req *dto.SignupRequest) error {
	if err := helper.ValidateUser(req); err != nil {
		return err
	}
 
 if req.Password != req.ConfirmPassword {
	return &helper.CustomError{
			Type:    helper.ErrBadRequest,
			Message: helper.ErrorResponse{Message: "Passwords do not match"},
	}
}

	err := repositories.CheckIfEmailExists(req.Email)
	if err != nil {
		return err
	}

	hashedPassword := helper.HashPassword(req.Password)
	createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	user := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		UserType:  req.UserType,
		CreatedAt: createdAt,
	}

	err = repositories.Signup(user)
	if err != nil {
		return err
	}

	return nil
}

func Login(email, password string) (*dto.UserResponse, string, string, error) {
	var user models.User

	user, err := repositories.GetUserByEmailForLogin(email)
	if err != nil {
		return nil, "", "", err
	}

	passwordIsValid, err := helper.VerifyPassword(password, user.Password)
	if !passwordIsValid {
		return nil, "", "", err
	}

	token, refreshToken, _ := helper.GenerateAllTokens(user.Name, user.Email, user.UserType, user.ID)

	err = helper.UpdateAllTokens(token, refreshToken, user.ID)
	if err != nil {
		return nil, "", "", err
	}
/* 
	user.Token = &token
	user.RefreshToken = &refreshToken */

	userResponse := &dto.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		UserType:     user.UserType,
		CreatedAt:    user.CreatedAt,
	}

	return userResponse, token, refreshToken, nil
}
