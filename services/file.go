package services

import (
	"idnatiya/go-auth/helpers"
	"idnatiya/go-auth/models"
	"idnatiya/go-auth/repositories"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileService struct {
	*repositories.FileRepository
}

func (service *FileService) Create(file *multipart.FileHeader, targetDir string, ctx *gin.Context) (*models.File, error) {
	extension, _ := helpers.Last[string](strings.Split(file.Filename, "."))
	fileName := uuid.New().String() + "." + extension

	uploadedFile := models.File{}
	uploadedFile.FileName = fileName
	uploadedFile.PathLocation = targetDir

	uploadDestination := targetDir + "/" + fileName

	ctx.SaveUploadedFile(file, "storage/app/public/"+uploadDestination)

	service.FileRepository.Create(&uploadedFile)

	return &uploadedFile, nil
}

func NewFileService(repo *repositories.FileRepository) *FileService {
	return &FileService{
		FileRepository: repo,
	}
}
