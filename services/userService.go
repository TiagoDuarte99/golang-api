package services

import (
	"github/tiagoduarte/golang-api/dto"
	helper "github/tiagoduarte/golang-api/helpers"
	repositories "github/tiagoduarte/golang-api/repositories"

	"github/tiagoduarte/golang-api/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsersWithPagination(ctx *gin.Context) ([]dto.UserResponse, error) {
	if err := helper.CheckUserType(ctx, "ADMIN"); err != nil {
		return nil, err
	}

	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	const recordPerPage = 10

	offset := (page - 1) * recordPerPage
	users, err := repositories.GetUsers(offset, recordPerPage)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByIDWithAuthorization(ctx *gin.Context, userId string) (*dto.UserResponse, error) {
	if err := helper.MatchUserTypeToUserId(ctx, userId); err != nil {
		return nil, err
	}

	userIDInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}

	user, err := repositories.GetUserByID(userIDInt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByIDWithAuthorizationForUpdate(ctx *gin.Context, userId string) (*models.User, error) {
	if err := helper.MatchUserTypeToUserId(ctx, userId); err != nil {
		return nil, err
	}

	userIDInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}

	user, err := repositories.GetUserCompleteByID(userIDInt)
	if err != nil {
		return nil, err
	}

	return user, nil
}


func UpdateUser(ctx *gin.Context, userId string, userUpdate dto.UpdateUserRequest) (*dto.UserResponse, error) {
	/* 	userIDInt, err := strconv.Atoi(userId)
	   	if err != nil {
	   		return nil, err
	   	} */

	user, err := GetUserByIDWithAuthorizationForUpdate(ctx, userId)
	if err != nil {
		return nil, err
	}

	if userUpdate.Name != "" {
		user.Name = userUpdate.Name
	}
	if userUpdate.Email != "" {
		err := repositories.CheckIfEmailExists(userUpdate.Email)
		if err != nil {
			return nil, err
		}
		user.Email = userUpdate.Email
	}

	if userUpdate.Password != "" {
		if userUpdate.Password != userUpdate.ConfirmPassword {
			passwordIsValid, err := helper.VerifyPassword(userUpdate.Password, userUpdate.ConfirmPassword)
			if !passwordIsValid {
				return nil, err
			}
		}

		user.Password = helper.HashPassword(userUpdate.Password)

	}

	updatedUser, err := repositories.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func DeleteUser(ctx *gin.Context, userId string) error {

	user, err := GetUserByIDWithAuthorization(ctx, userId)
	if err != nil {
		return err
	}

	err = repositories.DeleteUser(user)
	if err != nil {
		return err
	}

	return nil
}
