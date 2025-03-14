package main

import (
	"github/tiagoduarte/golang-api/database"
	routes "github/tiagoduarte/golang-api/routes"
	_ "github/tiagoduarte/golang-api/docs"

	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//buscar a porta em ENV
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
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

	router.StaticFile("/swagger.json", "./docs/swagger.json")

	// Rota Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
