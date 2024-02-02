package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	
	router.Use(gin.Logger())
	

	router.Run("localhost:8080")
}