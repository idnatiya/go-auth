package controllers

import (
	"idnatiya/go-auth/repositories"
	"idnatiya/go-auth/requests"

	"github.com/gin-gonic/gin"
)

type PermissionController struct{}

var permissionRepository = repositories.NewPermissionRepository()

func (ctrl PermissionController) CreatePermission(c *gin.Context) {
	var createPermissionRequest requests.CreatePermissionRequest
	if err := c.ShouldBindJSON(&createPermissionRequest); err != nil {
		c.Error(err)
		return
	}

	permissionRepository.SayHello()

	c.JSON(200, gin.H{
		"status": "success",
	})
}
