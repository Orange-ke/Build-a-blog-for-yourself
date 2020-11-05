package routes

import (
	"github.com/gin-gonic/gin"
	v1 "myblog/api/v1"
	"myblog/middleware"
	"myblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger(), middleware.Cors())
	r.Use(gin.Recovery())

	// 部署
	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static", "static/admin/static")
	r.Static("admin/favicon.ico", "static/admin/favicon.ico")

	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// 需要鉴权的接口
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.PUT("reset/:id", v1.ResetPassword)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		// 上传文件
		auth.POST("upload", v1.Upload)
	}

	// 不需要鉴权的api
	router := r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)
		// 分类模块的路由接口
		router.GET("category/:id", v1.GetCategory)
		router.GET("categories", v1.GetCategories)
		// 文章模块的路由接口
		router.GET("articles", v1.GetArticles) // 查询文章列表
		router.GET("article/:id", v1.GetArticle) // 查询单个文章
		// 登陆
		router.POST("login", v1.Login)
	}

	_ = r.Run(utils.HttpPort)
}