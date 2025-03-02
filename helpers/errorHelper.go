package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAuthorizationError(ctx *gin.Context, err error) bool {
	if err != nil {

		if err.Error() == "unauthorized" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized for this"})
			return true
		}

		if err.Error() == "notfound" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Requested resource not found"})
			return true
		}
		return true
	}
	return false
}