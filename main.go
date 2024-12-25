package main

import (
	. "onedollarback/Config"
	. "onedollarback/Model"
	. "onedollarback/Router"
)

func main() {

	InitMyDb()
	var err error
	err = MysqlDB.AutoMigrate(&User{})
	if err != nil {
		Logger.Error("数据库迁移失败")
	}

	r := InitRouter()

	err1 := r.Run("127.0.0.1:8080")
	if err1 != nil {
		Logger.Error("路由启动失败")

	}

}
