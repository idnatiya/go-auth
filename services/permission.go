package services

import (
	"idnatiya/go-auth/models"
	"idnatiya/go-auth/repositories"
	"idnatiya/go-auth/requests"
)

type PermissionService struct {
	*repositories.PermissionRepository
}

func (service *PermissionService) Get() (*[]models.Permission, error) {
	return service.PermissionRepository.Get()
}

func (service *PermissionService) FindByID(permissionID string) (*models.Permission, error) {
	return service.PermissionRepository.FindByID(permissionID)
}

func (service *PermissionService) Create(req *requests.PermissionRequest) (*models.Permission, error) {
	permission := models.Permission{
		PermissionName: req.PermissionName,
	}

	return service.PermissionRepository.Create(&permission)
}

func (service *PermissionService) Update(req *requests.PermissionRequest, permission *models.Permission) (*models.Permission, error) {
	permission.PermissionName = req.PermissionName

	return service.PermissionRepository.Update(permission)
}

func (service *PermissionService) Delete(permission *models.Permission) error {
	return service.PermissionRepository.Delete(permission)
}

func NewPermissionService(repo *repositories.PermissionRepository) *PermissionService {
	return &PermissionService{
		PermissionRepository: repo,
	}
}
