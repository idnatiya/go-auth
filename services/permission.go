package services

import (
	"idnatiya/go-auth/models"
	"idnatiya/go-auth/repositories"
	"idnatiya/go-auth/requests"
)

type PermissionService struct {
	*repositories.PermissionRepository
}

func (service PermissionService) Create(req *requests.CreatePermissionRequest) (*models.Permission, error) {
	permission := models.Permission{
		PermissionName: req.PermissionName,
	}

	return service.PermissionRepository.Create(&permission)
}

func NewPermissionService(repo *repositories.PermissionRepository) *PermissionService {
	return &PermissionService{
		PermissionRepository: repo,
	}
}
