package repositories

import (
	"idnatiya/go-auth/databases"
	"idnatiya/go-auth/models"
)

type PermissionRepository struct {
}

func (repo *PermissionRepository) Get() (*[]models.Permission, error) {
	var permissions []models.Permission
	err := databases.DB.Find(&permissions)
	if err.Error != nil {
		return nil, err.Error
	}

	return &permissions, err.Error
}

func (repo *PermissionRepository) FindByID(permissionID string) (*models.Permission, error) {
	var permission models.Permission
	query := databases.DB.First(&permission, "id = ?", permissionID)
	if query.Error != nil {
		return nil, query.Error
	}

	return &permission, nil
}

func (repo *PermissionRepository) Create(permission *models.Permission) (*models.Permission, error) {
	err := databases.DB.Save(permission)
	if err.Error != nil {
		return nil, err.Error
	}

	return permission, err.Error
}

func (repo *PermissionRepository) Update(permission *models.Permission) (*models.Permission, error) {
	err := databases.DB.Save(permission)
	if err.Error != nil {
		return nil, err.Error
	}

	return permission, err.Error
}

func (repo *PermissionRepository) Delete(permission *models.Permission) error {
	return databases.DB.Delete(&permission).Error
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}
