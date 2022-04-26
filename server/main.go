package main

import (
	"myapp/db"
	"myapp/route"
)

func main() {

	// 连接数据库
	db.Init()
	// 初始化路由，启动服务
	route.Run()
}
