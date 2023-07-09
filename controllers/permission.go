package controllers

import (
	"errors"
	"idnatiya/go-auth/requests"
	"idnatiya/go-auth/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PermissionController struct {
	*services.PermissionService
}

func (ctrl *PermissionController) ListPermission(context *gin.Context) {
	permissions, err := ctrl.PermissionService.Get()
	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(200, gin.H{
		"permissions": permissions,
	})
}

func (ctrl *PermissionController) ShowPermission(context *gin.Context) {
	permissionID := context.Param("id")
	permission, err := ctrl.PermissionService.FindByID(permissionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatusJSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.Error(err)
		return
	}

	context.JSON(200, gin.H{
		"permission": permission,
	})
}

func (ctrl *PermissionController) CreatePermission(context *gin.Context) {
	var createPermissionRequest requests.PermissionRequest
	if err := context.ShouldBindJSON(&createPermissionRequest); err != nil {
		context.Error(err)
		return
	}

	permission, err := ctrl.PermissionService.Create(&createPermissionRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			context.AbortWithStatusJSON(422, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(201, gin.H{
		"permission": permission,
	})
}

func (ctrl *PermissionController) UpdatePermission(context *gin.Context) {
	permission, err := ctrl.PermissionService.FindByID(context.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatusJSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.Error(err)
		return
	}

	var permissionRequest requests.PermissionRequest
	if err := context.ShouldBindJSON(&permissionRequest); err != nil {
		context.AbortWithError(400, err)
		return
	}

	updatedPermission, err := ctrl.PermissionService.Update(&permissionRequest, permission)
	if err != nil {
		context.AbortWithError(500, err)
		return
	}

	context.JSON(200, gin.H{
		"permission": updatedPermission,
	})
	return
}

func (ctrl *PermissionController) Delete(context *gin.Context) {
	permission, err := ctrl.PermissionService.FindByID(context.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatusJSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.Error(err)
		return
	}

	deletedErr := ctrl.PermissionService.Delete(permission)
	if deletedErr != nil {
		context.Error(deletedErr)
		return
	}

	context.JSON(204, nil)
}

func NewPermissionController(service *services.PermissionService) *PermissionController {
	return &PermissionController{
		PermissionService: service,
	}
}
