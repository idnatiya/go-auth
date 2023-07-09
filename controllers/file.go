package controllers

import (
	"idnatiya/go-auth/helpers"
	"idnatiya/go-auth/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	// Get file extension
	extension, _ := helpers.Last[string](strings.Split(file.Filename, "."))
	// Generate new file name by UUID
	newFileName := uuid.New().String() + "." + extension
	// Move file uploaded
	context.SaveUploadedFile(file, "storage/app/public/"+newFileName)

	context.JSON(200, gin.H{
		"message":     "Successfully uploaded file",
		"newFileName": newFileName,
	})
}

func NewFileController(service *services.FileService) *FileController {
	return &FileController{
		FileService: service,
	}
}
