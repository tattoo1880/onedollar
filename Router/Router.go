package Router

import (
	"github.com/gin-gonic/gin"
	. "onedollarback/Service"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	userapi := r.Group("/user")
	{
		userapi.POST("/create", CreateUser)
		userapi.POST("/get", GetUserByUsername)
		userapi.POST("/update", UpdateUser)
		userapi.POST("/delete", DeleteUserByID)
		userapi.GET("/getall", GetAllUsers)
		userapi.POST("/send", ReciveRabbitMQ)
	}

	return r

}
