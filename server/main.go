package main

import (
	"myapp/db"
	"myapp/global"
	"myapp/route"
)

func main() {

	// 连接数据库
	global.GB_DB = db.Init()
	// 初始化路由，启动服务
	route.Run()
}
