package dto

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}
