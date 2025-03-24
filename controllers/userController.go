package controllers

import (
	"github/tiagoduarte/golang-api/dto"
	helper "github/tiagoduarte/golang-api/helpers"
	"github/tiagoduarte/golang-api/services"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// @Summary Get users with pagination
// @Description This endpoint retrieves a list of users with pagination support.
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of users per page" default(10)
// @Success 200 {array} models.User "List of users"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 500 {object} helper.ErrorResponse "Internal Server Error"
// @Router /users [get]
func GetUsers(ctx *gin.Context) {
	// Verificação de autorização já vai para o serviço
	users, err := services.GetUsersWithPagination(ctx)
	if err != nil {
		helper.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// @Summary Get user by id
// @Description This endpoint retrieves a user by their ID.
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User "User data"
// @Failure 401 {object} helper.ErrorResponse "ErrUnauthorized"
// @Failure 404 {object} helper.ErrorResponse "ErrNotFound"
// @Failure 500 {object} helper.ErrorResponse "Internal Server Error"
// @Router /users/{id} [get]
func GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := services.GetUserByIDWithAuthorization(ctx, userId)
	if err != nil {
		helper.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary Update user
// @Description This endpoint allows a user to update their profile information.
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body dto.UpdateUserRequest true "User update data"
// @Success 200 {object} models.User "Updated user data"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 403 {object} helper.ErrorResponse "Forbidden"
// @Failure 404 {object} helper.ErrorResponse "User not found"
// @Failure 500 {object} helper.ErrorResponse "Internal Server Error"
// @Router /users/{id} [patch]
func UpdateUser(ctx *gin.Context) {
	userUpdate := dto.UpdateUserRequest{}

	if err := ctx.BindJSON(&userUpdate); err != nil {
		customErr := &helper.CustomError{
			Type:    helper.ErrBadRequest,
			Message: helper.ErrorResponse{Message: "Invalid request data: " + err.Error()},
		}
		helper.HandleError(ctx, customErr)
		return
	}

	userId := ctx.Param("id")

	updatedUser, err := services.UpdateUser(ctx, userId, userUpdate)
	if err != nil {
		helper.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

// @Summary Delete user
// @Description This endpoint allows an authenticated user to delete their account.
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string "User deleted successfully"
// @Failure 400 {object} helper.ErrorResponse "Bad Request"
// @Failure 401 {object} helper.ErrorResponse "Unauthorized"
// @Failure 403 {object} helper.ErrorResponse "Forbidden"
// @Failure 404 {object} helper.ErrorResponse "User not found"
// @Failure 500 {object} helper.ErrorResponse "Internal Server Error"
// @Router /users/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := services.DeleteUser(ctx, userId)
	if err != nil {
		helper.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
