package routes

import(
	controller "github/tiagoduarte/golang-api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine){
 	incomingRoutes.POST("users/signup", controller.Signup)
	incomingRoutes.POST("users/login", controller.Login)
}