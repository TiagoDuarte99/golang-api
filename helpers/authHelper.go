package helper

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckUserType(ctx *gin.Context, userTypeRole string) (err error) {
	userType := ctx.GetString("user_type")
	err = nil

	if userType != userTypeRole {
		err = errors.New("unauthorized to access this resource")
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
		err = errors.New("unauthorized to access this resource")
		return err
	}

	err = CheckUserType(ctx, userType)
	return err
}
