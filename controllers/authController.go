package controllers

import (
	"github/tiagoduarte/golang-api/dto"
	helper "github/tiagoduarte/golang-api/helpers"
	"github/tiagoduarte/golang-api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// @Summary Create a new user
// @Description This endpoint allows you to create a new user with the provided data.
// @Accept json
// @Produce json
// @Param user body dto.SignupRequest true "User data" // Certifique-se de usar o caminho correto para o tipo User
// @Success 200 {object} dto.SignupRequest
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 404 {object} helper.ErrorResponse "Not Found"
// @Failure 409 {object} helper.ErrorResponse "Conflit"
// @Router /users [post]
func Signup(ctx *gin.Context) {
	var user dto.SignupRequest

	if err := ctx.BindJSON(&user); err != nil {
		helper.HandleError(ctx, err)
		return
	}

	err := services.Signup(&user)
	log.Println(err)
	if err != nil {
		helper.HandleError(ctx, err)
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
