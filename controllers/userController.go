package controllers

import (
	helper "github/tiagoduarte/golang-api/helpers"
	"github/tiagoduarte/golang-api/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	"github/tiagoduarte/golang-api/database"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func Signup(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(user)

	if validationErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": validationErr.Error()})
		return
	}

	if err := database.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	user.Password = HashPassword(user.Password)
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "error to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})
}

/*
func Login()
*/

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
