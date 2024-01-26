package v1

import (
	"ginblog/model"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := model.UpLoadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
		"url":     url,
	})

}
