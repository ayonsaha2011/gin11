package main

import (
	"github.com/gin-gonic/gin"
	"github/ayonsaha2011/golang-gin-poc/controller"
	"github/ayonsaha2011/golang-gin-poc/middlewares"
	"github/ayonsaha2011/golang-gin-poc/service"
	"io"
	"net/http"
	"os"
)

var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput()  {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/videos", func (ctx *gin.Context)  {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func (ctx *gin.Context)  {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video input is Valid"})
		}

	})

	server.Run(":8080")
}