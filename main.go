package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/config"
	"github.com/joho/godotenv"
)


func init() {
	godotenv.Load()
	config.ConnectToDB()
}


func main() {
	router := gin.Default()
	
	router.Use(gin.Logger())
	

	router.Run("localhost:8080")
}