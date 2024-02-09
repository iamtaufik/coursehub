package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/config"
	"github.com/iamtaufik/coursehub/routes"
	"github.com/joho/godotenv"
)


func init() {
	godotenv.Load()
	config.ConnectToDB()
	config.SyncDatabase()
}


func Handler() {
	router := gin.Default()
	router.Use(gin.Logger())
	gin.SetMode(gin.ReleaseMode)

	routes.RegisterRoutes(router)

	router.Run("localhost:3000")
}