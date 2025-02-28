package controllers

import (
	"github/tiagoduarte/golang-api/dto"
	"github/tiagoduarte/golang-api/models"
	"github/tiagoduarte/golang-api/services"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func Signup(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.Signup(&user); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(ctx *gin.Context) {
	loginData := dto.LoginRequest{}

	if err := ctx.BindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, refreshToken, err := services.Login(loginData.Email, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":          user,
		"token":         token,
		"refresh_token": refreshToken,
	})
}
