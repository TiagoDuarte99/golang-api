package services

import (
	"errors"
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

func Login(email, password string) (models.User, string, string, error) {
	var user models.User

	user, err := repositories.GetUserByEmailForLogin(email)
	if err != nil {
		return user, "", "", err
	}

	passwordIsValid, msg := helper.VerifyPassword(password, user.Password)
	if !passwordIsValid {
		return user, "", "", errors.New(msg)
	}

	token, refreshToken, _ := helper.GenerateAllTokens(user.Name, user.Email, user.UserType, user.ID)

	err = helper.UpdateAllTokens(token, refreshToken, user.ID)
	if err != nil {
		return user, "", "", err
	}

	user.Token = &token
	user.RefreshToken = &refreshToken

	return user, token, refreshToken, nil
}
