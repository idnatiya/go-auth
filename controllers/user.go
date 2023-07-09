package controllers

import "github.com/gin-gonic/gin"

type UserController struct{}

func (ctrl UserController) CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "success",
	})
}
