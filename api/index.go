package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamtaufik/coursehub/config"
	"github.com/iamtaufik/coursehub/routes"
	"github.com/joho/godotenv"
)


var (
	app *gin.Engine
)

func init() {
	godotenv.Load()
	config.ConnectToDB()
	config.SyncDatabase()
	app = gin.Default()
	routes.RegisterRoutes(app)
}


func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}