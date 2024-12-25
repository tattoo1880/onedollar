package main

import (
	. "onedollarback/Config"
	. "onedollarback/Model"
	. "onedollarback/Router"
	. "onedollarback/Service"
)

func main() {
	InitMyDb()
	var err error
	err = MysqlDB.AutoMigrate(&User{})
	if err != nil {
		Logger.Error("数据库迁移失败")
	}

	// 2. 初始化RabbitMQ
	NewRabbitMQ()
	defer MyRabbitMQ.Close()

	// 3. go ConsumeService("onedollarback")
	go ConsumeService("onedollarback")

	r := InitRouter()
	// 1. 静态文件服务
	r.Static("/static", "./static")
	err1 := r.Run("127.0.0.1:8080")
	if err1 != nil {
		Logger.Error("路由启动失败")

	}

}
