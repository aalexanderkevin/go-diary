package main

import (
	"log"
	"os"
	"tap-talk/controllers"
	"tap-talk/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.Title = "Go-Diary"
	docs.SwaggerInfo.Description = "Personal Daily Diary Application for Go(lang)"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Initialize Router Server
	server := gin.Default()

	// Router Handlers
	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		register := apiRoutes.Group("/register")
		{
			register.POST("", controllers.InsertUser)
		}
		login := apiRoutes.Group("/login")
		{
			login.POST("", controllers.Login)
		}
		diary := apiRoutes.Group("/diary")
		{
			diary.GET("/:year/:quarter", controllers.GetDiary)
			diary.POST("", controllers.CreateDiary)

		}
		logout := apiRoutes.Group("/logout")
		{
			logout.DELETE("", controllers.Logout)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(server.Run(":" + os.Getenv("PORT")))

}
