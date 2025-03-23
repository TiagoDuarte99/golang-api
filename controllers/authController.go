package controllers

import (
	"github/tiagoduarte/golang-api/dto"
	helper "github/tiagoduarte/golang-api/helpers"
	"github/tiagoduarte/golang-api/services"

	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// @Summary Create a new user
// @Description This endpoint allows you to create a new user with the provided data.
// @Accept json
// @Produce json
// @Param user body dto.SignupRequest true "User data"
// @Success 200 {object} dto.SuccessMessage "User registered successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 404 {object} helper.ErrorResponse "Not Found"
// @Failure 409 {object} helper.ErrorResponse "Conflit"
// @Router /users [post]
func Signup(ctx *gin.Context) {
	var user dto.SignupRequest

	if err := ctx.BindJSON(&user); err != nil {
		customErr := &helper.CustomError{
			Type:    helper.ErrBadRequest,
			Message: helper.ErrorResponse{Message: "Invalid request data: " + err.Error()},
		}
		helper.HandleError(ctx, customErr)
		return
	}

	err := services.Signup(&user)

	if err != nil {
		helper.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.SuccessMessage{Message: "User registered successfully"})
}

// @Summary User Login
// @Description This endpoint allows an existing user to login by providing their email and password.
// @Accept json
// @Produce json
// @Param loginData body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 404 {object} helper.ErrorResponse "Not Found"
// @Failure 500 {object} helper.ErrorResponse "Internal Server Error"
// @Router /login [post]
func Login(ctx *gin.Context) {
	loginData := dto.LoginRequest{}

	if err := ctx.BindJSON(&loginData); err != nil {
		customErr := &helper.CustomError{
			Type:    helper.ErrBadRequest,
			Message: helper.ErrorResponse{Message: "Invalid request data: " + err.Error()},
		}
		helper.HandleError(ctx, customErr)
		return
	}

	user, token, refreshToken, err := services.Login(loginData.Email, loginData.Password)
	if err != nil {
		helper.HandleError(ctx, err)
		return
	}

	response := dto.LoginResponse{
		User:         user,         // Supondo que você tenha uma variável user com os dados do usuário
		Token:        token,        // Token gerado
		RefreshToken: refreshToken, // Token de refresh gerado
	}
	ctx.JSON(http.StatusOK, response)
}
