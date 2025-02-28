package services

import (
	helper "github/tiagoduarte/golang-api/helpers"
	repositories "github/tiagoduarte/golang-api/repositories"

	"github/tiagoduarte/golang-api/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsersWithPagination(ctx *gin.Context) ([]models.User, error) {
	// Checar tipo de usuário diretamente no service
	if err := helper.CheckUserType(ctx, "ADMIN"); err != nil {
		return nil, err
	}

	// Lógica de paginação movida para o service
	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	const recordPerPage = 4

	offset := (page - 1) * recordPerPage
	users, err := repositories.GetUsers(offset, recordPerPage)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByIDWithAuthorization(ctx *gin.Context, userId string) (*models.User, error) {
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
