package exceptions

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GlobalExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch err.Err.(type) {
			case validator.ValidationErrors:
				c.AbortWithStatusJSON(400, HandleValidationError(err))
				return
			}
		}
	}
}
