package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/config"
	"github.com/iamtaufik/coursehub/controllers"
	"github.com/joho/godotenv"
)


func init() {
	godotenv.Load()
	config.ConnectToDB()
	config.SyncDatabase()
}


func main() {
	router := gin.Default()
	
	router.Use(gin.Logger())

	router.Group("/api")
	{
		router.GET("/users", controllers.GetAllUsers)
		router.POST("/users", controllers.CreateUser)
	}

	router.Run("localhost:3000")
}