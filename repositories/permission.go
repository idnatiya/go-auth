package repositories

import "fmt"

type PermissionRepository interface {
	SayHello()
}

type permissionRepository struct {
}

func (repo permissionRepository) SayHello() {
	fmt.Println("Hello")
}

func NewPermissionRepository() PermissionRepository {
	var newPermissionRepository PermissionRepository
	newPermissionRepository = permissionRepository{}
	return newPermissionRepository
}
