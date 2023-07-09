//go:build wireinject
// +build wireinject

package providers

import (
	"idnatiya/go-auth/controllers"
	"idnatiya/go-auth/repositories"
	"idnatiya/go-auth/services"

	"github.com/google/wire"
)

func InitializePermission() *controllers.PermissionController {
	wire.Build(
		controllers.NewPermissionController,
		repositories.NewPermissionRepository,
		services.NewPermissionService,
	)
	return nil
}
