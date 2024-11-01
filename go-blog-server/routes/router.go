package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-blog-golang/server/api/v1"
	"my-blog-golang/server/middleware"
	"my-blog-golang/server/utils"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode) // 设置Gin启动模式，有debug、release两种

	//r := gin.Default()     // 创建路由
	r := gin.New()           // 创建路由
	r.Use(middleware.Log())  // 日志钩子
	r.Use(gin.Recovery())    // Gin框架默认的函数
	r.Use(middleware.Cors()) // 跨域

	// 加载图片资源目录，用于存储用户上传的缩略图头像等
	// 此处建议存储到阿里云OSS或者其他的云上，然后开启CDN加快访问速度
	r.StaticFS(fmt.Sprintf("/%s", utils.ImgPath), http.Dir(utils.ImgPath))

	auth := r.Group("api/v1")       // 设置路由组
	auth.Use(middleware.JwtToken()) // 鉴权中间件钩子
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.POST("user/add", v1.AddUser)

		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)

		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

		// 上传文件
		auth.POST("upload", v1.UpLoad)
	}

	router := r.Group("api/v1") // 设置路由组
	{
		// 用户模块的路由接口
		//router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("user/:id", v1.GetUserInfo) // 查询单个用户

		// 分类模块的路由接口
		router.GET("category", v1.GetCate)

		// 文章模块的路由接口
		router.GET("article", v1.GetArt)              // 获取所有文章列表
		router.GET("article/list/:id", v1.GetCateArt) // 查询分类下的所有文章
		router.GET("article/info/:id", v1.GetArtInfo) // 查询单个文章信息

		// 登录
		router.POST("login", v1.Login)
		//router.POST("upload", v1.UpLoad)
	}

	r.Run(utils.HttpPort) // 启动程序
}
