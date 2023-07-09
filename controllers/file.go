package controllers

import (
	"idnatiya/go-auth/services"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	*services.FileService
}

func (ctrl *FileController) Create(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		context.AbortWithStatusJSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}

	uploadedFile, err := ctrl.FileService.Create(file, "", context)
	if err != nil {
		context.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
	}

	context.JSON(200, gin.H{
		"message": "Successfully uploaded file",
		"file":    uploadedFile,
	})
}

func NewFileController(service *services.FileService) *FileController {
	return &FileController{
		FileService: service,
	}
}
