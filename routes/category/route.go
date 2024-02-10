package categoryRoute

import (
	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/controllers"
)

func RegisterCategoryRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetCategories)

}