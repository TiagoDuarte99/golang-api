package main

import (
	"github/tiagoduarte/golang-api/database"
	routes "github/tiagoduarte/golang-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main(){
	//buscar a porta em ENV
	port:= os.Getenv("PORT")

	if port==""{
		port= "8000"
	}
	
// Inicializar o banco de dados
	database.InitDB()

	//Iniciar o router com o gin
	router := gin.New()
	router.Use(gin.Logger())

	//importar as Rotas 
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.TeamsRoutes(router)


	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-2"})
	})
	
	router.Run(":" + port)
}