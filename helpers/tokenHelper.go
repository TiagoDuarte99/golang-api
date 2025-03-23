package helper

import (
	"github/tiagoduarte/golang-api/database"
	"github/tiagoduarte/golang-api/dto"
	"github/tiagoduarte/golang-api/models"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(name, email, userType string, id int) (string, string, error) {

	claims := &dto.TokenClaims{
		ID:       id,
		Name:     name,
		Email:    email,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Expira em 24 horas
		},
	}

	// Criando o refresh token (com duração maior)
	refreshClaims := &dto.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(), // Expira em 7 dias
		},
	}

	// Assinando os tokens
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func ValidateToken(signedToken string) (*dto.TokenClaims, string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&dto.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return nil, "Erro ao validar o token"
	}

	claims, ok := token.Claims.(*dto.TokenClaims)
	if !ok {
		return nil, "Token inválido"
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, "Token expirado"
	}

	return claims, ""
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userID int) error {
	err := database.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"token":         signedToken,
			"refresh_token": signedRefreshToken,
		}).Error

	return err
}
