# 使用方式
## 请求方式
url/api/v1/interface


## 管理接口
auth.PUT("user/:id", v1.EditUser)

auth.DELETE("user/:id", v1.DeleteUser)

auth.POST("category/add", v1.AddCategory)

auth.PUT("category/:id", v1.EditCategory)

auth.DELETE("category/:id", v1.DeleteCategory)

auth.POST("article/add", v1.AddArticle)

auth.PUT("article/:id", v1.EditArticle)

auth.DELETE("article/:id", v1.DeleteArticle)

auth.POST("upload", v1.Upload)

## 常规接口
router.POST("user/add", v1.AddUser)

router.GET("users", v1.GetUsers)

router.GET("categories", v1.GetCategories)

router.GET("articles", v1.GetArticles)

router.GET("article/category/:id", v1.GetCateArt)

router.GET("article/info/:id", v1.GetArticle)

router.POST("login", v1.Login)
