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

func (ctrl *PermissionController) CreatePermission(context *gin.Context) {
	var createPermissionRequest requests.CreatePermissionRequest
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

func NewPermissionController(service *services.PermissionService) *PermissionController {
	return &PermissionController{
		PermissionService: service,
	}
}
