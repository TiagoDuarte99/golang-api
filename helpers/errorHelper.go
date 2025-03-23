package helper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type CustomError struct {
Type    error
Message ErrorResponse
}


func (e *CustomError) Error() string {
	return  e.Message.Message
}

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")
	ErrBadRequest =  errors.New("bad request")
	ErrConflict     = errors.New("conflict") 
)

func HandleError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}
	/* log.Println(err) */

	if customErr, ok := err.(*CustomError); ok {
		switch customErr.Type {
		case ErrUnauthorized:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": customErr.Message.Message})
		case ErrNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": customErr.Message.Message})
		case ErrBadRequest:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": customErr.Message.Message})
		case ErrConflict:
			ctx.JSON(http.StatusConflict, gin.H{"error": customErr.Message.Message})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
	}

	ctx.Abort()
}

