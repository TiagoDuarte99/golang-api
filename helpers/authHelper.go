package helper

import (
	
	"github/tiagoduarte/golang-api/dto"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func CheckUserType(ctx *gin.Context, userTypeRole string) error {
	userType := ctx.GetString("user_type")
	if userType != userTypeRole {
		return &CustomError{
			Type:    ErrUnauthorized,
			Message: ErrorResponse{
        Message: "User type " + userType + " is not authorized. Required role: " + userTypeRole,
    },
		}
	}
	return nil
}

func MatchUserTypeToUserId(ctx *gin.Context, userId string) (err error) {
	userType := ctx.GetString("user_type")
	//Recebo o id como int e converto para string
	uid := strconv.Itoa(ctx.GetInt("id"))

	err = nil

	if userType == "USER" && uid != userId {
		return &CustomError{
			Type:    ErrUnauthorized,
			Message: ErrorResponse{
				"User with id " + uid + " is not authorized to access this resource. Required userId: " + userId,
			},
		}
	}


	return CheckUserType(ctx, userType)
}

func ValidateUser(user *dto.SignupRequest) error {

	if validationErr := validate.Struct(user); validationErr != nil {
			var errMessages []string
			for _, err := range validationErr.(validator.ValidationErrors) {
					errMessages = append(errMessages, err.Field()+" is invalid: "+err.Tag())
			}
			return &CustomError{
				Type:    ErrBadRequest,
				Message: ErrorResponse{Message: strings.Join(errMessages, ", ")},
			}
	}

	return nil
}
