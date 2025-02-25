package controllers

import (
	helper "github/tiagoduarte/golang-api/helpers"
	"github/tiagoduarte/golang-api/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"

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

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "email or password is incorrect"
		check = false
	}
	return check, msg
}

func Signup(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(user)

	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	if err := database.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	user.Password = HashPassword(user.Password)
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(ctx *gin.Context) {
	var user models.User
	var userFound models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("email = ?", user.Email).First(&userFound).Error; err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "email or password is incorrect"})
		return
	}

	passwordIsValid, msg := VerifyPassword(user.Password, userFound.Password)
	if !passwordIsValid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := helper.GenerateAllTokens(userFound.Name, userFound.Email, userFound.UserType, userFound.ID)

	helper.UpdateAllTokens(token, refreshToken, userFound.ID)

	err := database.DB.Where("id = ?", userFound.ID).First(&userFound).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userFound)
}

func GetUsers(ctx *gin.Context) {
	var users []models.User
	if err := helper.CheckUserType(ctx, "ADMIN"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recordPerPage := 4

	pageStr := ctx.DefaultQuery("page", "1")
	page, err1 := strconv.Atoi(pageStr)
	if err1 != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * recordPerPage

	if err := database.DB.Offset(offset).Limit(recordPerPage).Find(&users).Error; err != nil {
    ctx.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
    return
}

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
