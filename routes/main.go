package routes

import (
	"github.com/gin-gonic/gin"
	usersRoute "github.com/iamtaufik/coursehub/routes/users"
)

// RegisterRoutes is a function to register all routes

func RegisterRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		usersRoute.RegisterUserRoutes(apiGroup.Group("/users"))
	}
}