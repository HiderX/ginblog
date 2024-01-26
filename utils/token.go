package utils

import (
	"github.com/gin-gonic/gin"
)

func ResponseWithMsg(c *gin.Context, err int) {
	c.JSON(200, gin.H{
		"status":  err,
		"message": GetErrMsg(err),
	})
}
