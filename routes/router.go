package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/api/v1"
	"go-blog/middleware"
	"go-blog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	server := gin.New()
	server.Use(middleware.Logger())
	server.Use(gin.Recovery())
	server.Use(middleware.Cors())

	// 有权限的路由
	auth := server.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		auth.DELETE("/user/:id", v1.DeleteUser)
		auth.PUT("/user/:id", v1.EditUser)
		auth.GET("/users", v1.GetUsers)

		auth.POST("/category", v1.CreateCategory)
		auth.DELETE("/category/:id", v1.DeleteCategory)
		auth.PUT("/category/:id", v1.EditCategory)

		auth.POST("/article", v1.CreateArticle)
		auth.DELETE("/article/:id", v1.DeleteArticle)
		auth.PUT("/article/:id", v1.EditArticle)

		auth.POST("/upload", v1.Upload)
	}

	// 无权限路由
	routerV1 := server.Group("api/v1")
	{
		//登录路由
		routerV1.POST("/login", v1.Login)
		//用户模块路由
		routerV1.POST("/user", v1.CreateUser)
		//分类模块路由
		routerV1.GET("/categories", v1.GetCategory)
		//文章模块路由
		routerV1.GET("/article/:id", v1.GetSingleArticle)
		routerV1.GET("/articles", v1.GetArticles)
		routerV1.GET("/articles/:cid", v1.GetArticlesByCategory)
	}

	_ = server.Run(utils.HttpPort)
}
