package api

import (
	"adv/controllers"
	"adv/middleware"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// created id struct for path validation

func logging() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}

func GetRoutes() *gin.Engine {
	logging()
	r := gin.Default()
	r.GET("/advertisements", controllers.GetAllAdvertisements)
	r.GET("/advertisements/:id", middleware.ValidateIdParam, controllers.GetAdvertisementByID)
	r.POST("/advertisements", controllers.PostAdvertisement)
	auth := r.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/signin", controllers.SignIn)

	}
	return r
}
