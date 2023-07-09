package repositories

import (
	"idnatiya/go-auth/databases"
	"idnatiya/go-auth/models"
)

type PermissionRepository struct {
}

func (repo PermissionRepository) Create(permission *models.Permission) (*models.Permission, error) {
	err := databases.DB.Save(permission)
	if err.Error != nil {
		return nil, err.Error
	}

	return permission, err.Error
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}
