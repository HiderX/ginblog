package v1

import (
	"ginblog/model"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code := model.CreateArt(&data)
	c.JSON(200, gin.H{
		"status":  200,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}

func GetCateArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetCateArt(id, pageSize, pageNum)
	c.JSON(200, gin.H{
		"status":  200,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(200, gin.H{
		"status":  200,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}

func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetArt(pageSize, pageNum)
	c.JSON(200, gin.H{
		"status":  200,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}

func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.EditArt(id, &data)
	c.JSON(200, gin.H{
		"status":  200,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArt(id)
	c.JSON(200, gin.H{
		"status":  200,
		"message": utils.GetErrMsg(code),
	})
}
