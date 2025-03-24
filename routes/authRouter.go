package routes

import(
	controller "github/tiagoduarte/golang-api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine){
 	incomingRoutes.POST("/signup", controller.Signup)
	incomingRoutes.POST("/login", controller.Login)
}