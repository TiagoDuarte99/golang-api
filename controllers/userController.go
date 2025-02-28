package controllers

import (
	"github/tiagoduarte/golang-api/services"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func GetUsers(ctx *gin.Context) {
	// Verificação de autorização já vai para o serviço
	users, err := services.GetUsersWithPagination(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := services.GetUserByIDWithAuthorization(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}