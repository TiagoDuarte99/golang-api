package routes

import(
	controller "github/tiagoduarte/golang-api/controllers"
	middleware "github/tiagoduarte/golang-api/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.AuthMiddleware())
	incomingRoutes.GET("/users", controller.GetUsers) 
	incomingRoutes.GET("/users/:id", controller.GetUser)
	incomingRoutes.PATCH("/users/:id", controller.UpdateUser)
	incomingRoutes.DELETE("/users/:id", controller.DeleteUser)

}