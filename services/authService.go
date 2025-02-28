package services

import (
	"errors"
	helper "github/tiagoduarte/golang-api/helpers"
	repositories "github/tiagoduarte/golang-api/repositories"

	"github/tiagoduarte/golang-api/models"

	"time"

)

func Signup(user *models.User) error {
	if err := helper.ValidateUser(user); err != nil {
		return err
	}

	err := repositories.CheckIfEmailExists(user.Email)
	if err != nil {
		return err
	}

	user.Password = helper.HashPassword(user.Password)
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

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
	helper.UpdateAllTokens(token, refreshToken, user.ID)

	return user, token, refreshToken, nil
}
