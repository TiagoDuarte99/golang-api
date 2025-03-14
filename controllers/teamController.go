package controllers

import (
	/* 	"github/tiagoduarte/golang-api/dto" */
	/* 	"github/tiagoduarte/golang-api/models" */
	helper "github/tiagoduarte/golang-api/helpers"
	"github/tiagoduarte/golang-api/services"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func GetTeams(ctx *gin.Context) {
	users, err := services.GetTeamsWithPagination(ctx)
	if err != nil {
		helper.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func GetTeam(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := services.GetUserByIDWithAuthorization(ctx, userId)
	if err != nil {
		helper.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

/*
func UpdateTeam(ctx *gin.Context) {
	userUpdate := dto.UpdateUserRequest{}

	if err := ctx.BindJSON(&userUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	userId := ctx.Param("id")

	updatedUser, err := services.UpdateUser(ctx, userId, userUpdate)
	if err != nil {
		if helper.HandleError(ctx, err) {
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func DeleteTeam(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := services.DeleteUser(ctx, userId)
	if err != nil {
		if helper.HandleError(ctx, err) {
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
} */
