package services

import "idnatiya/go-auth/repositories"

type FileService struct {
	*repositories.FileRepository
}

func NewFileService(repo *repositories.FileRepository) *FileService {
	return &FileService{
		FileRepository: repo,
	}
}
