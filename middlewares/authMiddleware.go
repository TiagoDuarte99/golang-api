package middleware

import (
	"net/http"

	helper "github/tiagoduarte/golang-api/helpers"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if len(tokenString) > len("Bearer ") && tokenString[:len("Bearer ")] == "Bearer " {
			tokenString = tokenString[len("Bearer "):]
		} else {
			// Se o token não estiver presente, retorna erro
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			ctx.Abort()
			return
		}

		// Validar o token
		claims, msg := helper.ValidateToken(tokenString)
		if msg != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			ctx.Abort()
			return
		}

		// Guardando os dados do utilizador na requisição
		ctx.Set("id", claims.ID)
		ctx.Set("email", claims.Email)
		ctx.Set("name", claims.Name)
		ctx.Set("user_type", claims.UserType)
		ctx.Next()
	}
}
