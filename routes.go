package main

import "github.com/gin-gonic/gin"

func (app *application) Routes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")

	imageRoutes := v1.Group("/image")
	{
		imageRoutes.POST("/upload-single", app.UploadImageHandler)
		imageRoutes.POST("/upload-multiple", app.UploadMultipleFilesHandler)
	}

	return router
}
