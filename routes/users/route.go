package usersRoute

import (
	"github.com/iamtaufik/coursehub/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetAllUsers)
	router.POST("/", controllers.CreateUser)
}