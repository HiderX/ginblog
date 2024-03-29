package v1

import (
	"ginblog/model"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

//func UserExist(c *gin.Context) {
//	username := c.Query("username")
//	code := model.CheckUser(username)
//	c.JSON(200, gin.H{
//		"status":  code,
//		"message": utils.GetErrMsg(code),
//	})
//}

func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	msg, code := utils.Validate(&data)
	if code != utils.SUCCESS {
		c.JSON(200, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = model.CreateUser(&data)
	c.JSON(200, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetUsers(pageSize, pageNum)
	code := utils.SUCCESS
	c.JSON(200, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": utils.GetErrMsg(code),
	})
}

func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code := model.EditUser(id, &data)
	c.JSON(200, gin.H{
		"status":  code,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(200, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
