package repositories

import (
	"idnatiya/go-auth/databases"
	"idnatiya/go-auth/models"
)

type FileRepository struct {
}

func (repo *FileRepository) Create(file *models.File) (*models.File, error) {
	query := databases.DB.Save(&file)
	if query.Error != nil {
		return nil, query.Error
	}

	return file, nil
}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}
