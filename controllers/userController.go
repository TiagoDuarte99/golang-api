package controllers

import (
	"github/tiagoduarte/golang-api/dto"
/* 	"github/tiagoduarte/golang-api/models" */
	"github/tiagoduarte/golang-api/services"
	"net/http"
helper "github/tiagoduarte/golang-api/helpers"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func GetUsers(ctx *gin.Context) {
	// Verificação de autorização já vai para o serviço
	users, err := services.GetUsersWithPagination(ctx)
	if err != nil {
		if helper.HandleAuthorizationError(ctx, err) {
			return
		}
	}

	ctx.JSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := services.GetUserByIDWithAuthorization(ctx, userId)
	if err != nil {
		if helper.HandleAuthorizationError(ctx, err) {
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
		if helper.HandleAuthorizationError(ctx, err) {
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := services.DeleteUser(ctx, userId)
	if err != nil {
		if helper.HandleAuthorizationError(ctx, err) {
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}