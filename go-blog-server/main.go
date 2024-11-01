package main

import (
	"my-blog-golang/server/model"
	"my-blog-golang/server/routes"
	"my-blog-golang/server/utils"
)

func main() {
	// 创建一些目录资源
	utils.InitDir()

	// 初始化数据库
	model.InitDb()

	// 引用路由
	routes.InitRouter()
}
