package helper

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Name     string
	Email    string
	Password string
	UserType string
	jwt.StandardClaims 
}


