package main

import (
	"myblog/model"
	"myblog/routes"
)

func main() {
	// 初始化数据库
	db := model.InitDb()
	defer db.Close()

	// 启动路由
	routes.InitRouter()
}
