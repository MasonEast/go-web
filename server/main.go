package main

import (
	"myapp/db"
	"myapp/global"
	"myapp/kafka"
	"myapp/route"
)

func main() {

	// 初始化kafka消费者
	// consumer := new(kafka.Consumer)
	// go consumer.Start()
	go kafka.ConsumerGroup(global.TOPIC, "ConsumerGroupID", "C1")
	// 连接数据库
	global.GB_DB = db.Init()
	// 初始化路由，启动服务
	route.Run()
	
}
