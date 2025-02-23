package controllers

import (
	helper "github/tiagoduarte/golang-api/helpers"
	"github/tiagoduarte/golang-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	"github/tiagoduarte/golang-api/database"

	_ "github.com/lib/pq"
)

var validate = validator.New()
/* 
func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers() */

func GetUsers(ctx *gin.Context) {
	var users []models.User
/* 	if err := helper.CheckUserType(ctx, "ADMIN"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	} */

	if err := database.DB.Find(&users).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	//Falta paginação
	ctx.JSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {
	var user models.User
	userId := ctx.Param("id") 

	if err := helper.MatchUserTypeToUserId(ctx, userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.First(&user, userId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

