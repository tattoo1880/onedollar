package Config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func InitMyDb() {
	//1. 使用gorm连接数据库
	dsn := "root:qwerty7788421@tcp(127.0.0.1:3306)/onedollar?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Error("连接数据库失败，检查配置文件")
		panic("failed to connect database")
	}
	//2. 日志记录连接成功
	Logger.Info("连接数据库成功")
}
