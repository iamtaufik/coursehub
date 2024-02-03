package authenticationRoute

import (
	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/controllers"
	"github.com/iamtaufik/coursehub/middlewares"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/whoami",middlewares.VerifyToken, controllers.Whoami)
}