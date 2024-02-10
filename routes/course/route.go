package courseRoute

import (
	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/controllers"
	"github.com/iamtaufik/coursehub/middlewares"
)

func RegisterCourseRoutes(router *gin.RouterGroup) {
	router.POST("/", middlewares.VerifyToken, middlewares.IsAdmin, controllers.CreateCourse)
	router.GET("/join/:id", middlewares.VerifyToken, controllers.JoinCourse)
	router.GET("/me", middlewares.VerifyToken, controllers.MyCourses)
	router.GET("/", controllers.GetCourses)
	router.GET("/:id", controllers.GetCourse)
	// router.PUT("/course/:id", middlewares.VerifyToken, controllers.UpdateCourse)
	// router.DELETE("/course/:id", middlewares.VerifyToken, controllers.DeleteCourse)
}