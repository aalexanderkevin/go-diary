package routes

import (
	"log"
	"os"
	"tap-talk/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	// Initialize Router
	router := gin.Default()

	// Router Handlers
	router.POST("/register", controllers.InsertUser)
	router.POST("/login", controllers.Login)
	router.GET("/diary/:year/:quarter", controllers.GetDiary)
	router.POST("/diary", controllers.CreateDiary)
	router.DELETE("/logout", controllers.Logout)

	log.Fatal(router.Run(":" + os.Getenv("PORT")))

}
