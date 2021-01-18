package main

import (
	"log"
	"os"
	"tap-talk/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Router Server
	server := gin.Default()

	// Router Handlers
	server.POST("/register", controllers.InsertUser)
	server.POST("/login", controllers.Login)
	server.GET("/diary/:year/:quarter", controllers.GetDiary)
	server.POST("/diary", controllers.CreateDiary)
	server.DELETE("/logout", controllers.Logout)

	log.Fatal(server.Run(":" + os.Getenv("PORT")))

}
