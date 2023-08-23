package api

import (
	"adv/controllers"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func logging() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}

func GetRoutes() *gin.Engine {
	logging()
	r := gin.Default()
	r.GET("/advertisements", controllers.GetAllAdvertisements)
	r.GET("/advertisements/:id", controllers.GetAdvertisementByID)
	r.POST("/advertisements", controllers.PostAdvertisement)
	return r
}
