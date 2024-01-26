package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code := model.CheckLogin(data.Username, data.Password)
	if code == utils.SUCCESS {
		token, code := middleware.SetToken(data.Username, data.Password)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": utils.GetErrMsg(code),
			"token":   token,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  code,
			"message": utils.GetErrMsg(code),
		})
	}
}
