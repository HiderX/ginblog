package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{

		//auth.POST("user/add", v1.AddUser)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		auth.POST("upload", v1.Upload)
	}
	router := r.Group("api/v1")
	{
		// User module routing
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		//router.GET("user/:username", v1.UserExist)
		// Category module routing
		router.GET("categories", v1.GetCategories)
		// Article module routing
		router.GET("articles", v1.GetArticles)
		router.GET("article/category/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArticle)
		router.POST("login", v1.Login)
	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
