package services

import (

	"github/tiagoduarte/golang-api/dto"

	repositories "github/tiagoduarte/golang-api/repositories"

/* 	"github/tiagoduarte/golang-api/models" */

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTeamsWithPagination(ctx *gin.Context) ([]dto.TeamDTO, error) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	const recordPerPage = 4

	offset := (page - 1) * recordPerPage
	teams, err := repositories.GetTeams(offset, recordPerPage)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

/* func GetUserByIDWithAuthorization(ctx *gin.Context, userId string) (*models.User, error) {
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

	return &user, nil
}

func UpdateUser(ctx *gin.Context, userId string, userUpdate dto.UpdateUserRequest) (*models.User, error) {


	user, err := GetUserByIDWithAuthorization(ctx, userId)
	if err != nil {
		return nil, err
	}

	if userUpdate.Name != "" {
		user.Name = userUpdate.Name
	}
	if userUpdate.Email != "" {
		user.Email = userUpdate.Email
	}
	if userUpdate.Password != "" {
		if userUpdate.Password != userUpdate.ConfirmPassword {
			passwordIsValid, err := helper.VerifyPassword(userUpdate.Password, userUpdate.ConfirmPassword)
			if !passwordIsValid {
				return user, errors.New(err)
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
 */