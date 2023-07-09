package routes

import (
	"idnatiya/go-auth/providers"

	"github.com/gin-gonic/gin"
)

var permission = providers.InitializePermission()

func DefineWebRoute(c *gin.Engine) {
	v1 := c.Group("/v1")
	{
		v1.POST("/permission", permission.CreatePermission)
	}
}
