package controllers

import (
	"github/tiagoduarte/golang-api/dto"
/* 	"github/tiagoduarte/golang-api/models" */
	"github/tiagoduarte/golang-api/services"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func GetUsers(ctx *gin.Context) {
	// Verificação de autorização já vai para o serviço
	users, err := services.GetUsersWithPagination(ctx)
	if err != nil {
		if HandleAuthorizationError(ctx, err) {
			return
		}
	}

	ctx.JSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := services.GetUserByIDWithAuthorization(ctx, userId)
	if err != nil {
		if HandleAuthorizationError(ctx, err) {
			return
		}
	}

	ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context) {
	userUpdate := dto.UpdateUserRequest{}

	if err := ctx.BindJSON(&userUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}


	userId := ctx.Param("id")

	updatedUser, err := services.UpdateUser(ctx, userId, userUpdate)
	if err != nil {
		if HandleAuthorizationError(ctx, err) {
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func HandleAuthorizationError(ctx *gin.Context, err error) bool {
	if err != nil {

		if err.Error() == "unauthorized" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized for this"})
			return true
		}

		if err.Error() == "notfound" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return true
		}
		return true
	}
	return false
}