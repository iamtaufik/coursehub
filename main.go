package main

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

func main() {
	app := gin.Default()
	routes.RegisterRoutes(app)
	app.Run("localhost:3000")
}