package helpers

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"errors": err,
	})
}
