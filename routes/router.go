package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		// User module routing
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		//router.GET("user/:username", v1.UserExist)
		// Category module routing
		router.POST("category/add", v1.AddCategory)
		router.GET("categories", v1.GetCategories)
		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)
		// Article module routing
		router.POST("article/add", v1.AddArticle)
		router.GET("articles", v1.GetArticles)
		router.PUT("article/:id", v1.EditArticle)
		router.DELETE("article/:id", v1.DeleteArticle)
	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
