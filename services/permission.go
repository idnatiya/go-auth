package services

import (
	"idnatiya/go-auth/repositories"
	"idnatiya/go-auth/requests"
)

var permissionRepository = new(repositories.PermissionRepository)

type PermissionService interface {
	Create(req *requests.CreatePermissionRequest)
}

type permissionService struct{}

func (service permissionService) Create(req *requests.CreatePermissionRequest) {

}

func NewPermissionService() PermissionService {
	var newPermissionService PermissionService
	newPermissionService = permissionService{}
	return newPermissionService
}
