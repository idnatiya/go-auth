package routes

import (
	"idnatiya/go-auth/controllers"

	"github.com/gin-gonic/gin"
)

var user = new(controllers.UserController)
var permission = new(controllers.PermissionController)

func DefineWebRoute(c *gin.Engine) {
	v1 := c.Group("/v1")
	{
		v1.POST("/user", user.CreateUser)

		v1.POST("/permission", permission.CreatePermission)
	}
}
