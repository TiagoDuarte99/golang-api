package helper

import (
	"errors"
	"github/tiagoduarte/golang-api/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)
var validate = validator.New()

func CheckUserType(ctx *gin.Context, userTypeRole string) (err error) {
	userType := ctx.GetString("user_type")
	err = nil

	if userType != userTypeRole {
		err = errors.New("unauthorized")
		return err
	}
	return err
}

func MatchUserTypeToUserId(ctx *gin.Context, userId string) (err error) {
	userType := ctx.GetString("user_type")
	//Recebo o id como int e converto para string
	uid := strconv.Itoa(ctx.GetInt("id"))
	
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized")
		return err
	}

	err = CheckUserType(ctx, userType)
	return err
}

func ValidateUser(user *models.User) error {
	if validationErr := validate.Struct(user); validationErr != nil {
		var errMessages []string
		for _, err := range validationErr.(validator.ValidationErrors) {
			errMessages = append(errMessages, err.Field()+" is invalid: "+err.Tag())
		}
		return errors.New(strings.Join(errMessages, ", "))
	}

	return nil
}
