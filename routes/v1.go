package routes

import (
	"idnatiya/go-auth/providers"

	"github.com/gin-gonic/gin"
)

var permission = providers.InitializePermission()
var filesystem = providers.InitializeFilesystem()

func DefineWebRoute(c *gin.Engine) {
	v1 := c.Group("/v1")
	{
		v1.GET("/permissions", permission.ListPermission)
		v1.GET("/permission/:id/show", permission.ShowPermission)
		v1.POST("/permission", permission.CreatePermission)
		v1.PUT("/permission/:id", permission.UpdatePermission)
		v1.DELETE("/permission/:id", permission.Delete)

		v1.POST("/file", filesystem.Create)
	}
}
