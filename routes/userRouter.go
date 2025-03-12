package routes

import (
	controller "github/tiagoduarte/golang-api/controllers"
	middleware "github/tiagoduarte/golang-api/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// Criar um grupo de rotas protegido pelo middleware
	protectedRoutes := incomingRoutes.Group("/users")
	protectedRoutes.Use(middleware.AuthMiddleware())

	protectedRoutes.GET("/", controller.GetUsers)
	protectedRoutes.GET("/:id", controller.GetUser)
	protectedRoutes.PATCH("/:id", controller.UpdateUser)
	protectedRoutes.DELETE("/:id", controller.DeleteUser)
}
