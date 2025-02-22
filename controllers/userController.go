package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go.playground/validator/v10"
	helper "github/tiagoduarte/golang-api/helpers"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/akhil/golang-jwt-project/database"
	"github.com/akhil/golang-jwt-project/models"
	"golang.org/x/crypto/bcrypt"

	"database/sql"
	_ "github.com/lib/pq"
)


var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func