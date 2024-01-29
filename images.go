package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

func (app *application) UploadImageHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	extension := filepath.Ext(file.Filename)

	newFileName := uuid.New().String() + extension

	destinationPath := filepath.Join("uploads", newFileName)

	if err := c.SaveUploadedFile(file, destinationPath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File has been successfully uploaded"})
}

func (app *application) UploadMultipleFilesHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error parsing multipart form",
		})
		return
	}

	files := form.File["files"]

	if len(files) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No files are received",
		})
		return
	}

	for _, file := range files {
		extension := filepath.Ext(file.Filename)

		newFileName := uuid.New().String() + extension

		destinationPath := filepath.Join("uploads", newFileName)

		if err := c.SaveUploadedFile(file, destinationPath); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save one or more files",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "All files have been successfully uploaded"})
}
