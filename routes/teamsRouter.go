package routes

import (
	controller "github/tiagoduarte/golang-api/controllers"
	"log"
	/* 	middleware "github/tiagoduarte/golang-api/middlewares" */

	"github.com/gin-gonic/gin"
)

func TeamsRoutes(incomingRoutes *gin.Engine) {
	// Rotas públicas (sem middleware)
	log.Println(incomingRoutes)
	incomingRoutes.GET("/teams", controller.GetTeams)
	incomingRoutes.GET("/teams/:id", controller.GetTeam)
/* 
	// Grupo de rotas protegidas pelo middleware
	protectedRoutes := incomingRoutes.Group("/teams")
	protectedRoutes.Use(middleware.AuthMiddleware())
	protectedRoutes.PATCH("/:id", controller.UpdateTeam)
	protectedRoutes.DELETE("/:id", controller.DeleteTeam) */
}
