package Service

import (
	"github.com/gin-gonic/gin"
	. "onedollarback/Config"
	"onedollarback/Model"
)

func CreateUser(c *gin.Context) {

	var user Model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		Logger.Error("参数绑定失败")
		c.JSON(200, gin.H{
			"message": "参数绑定失败",
		})
	}
	if err := user.CreateUser(); err != nil {
		Logger.Error("创建用户失败")
		c.JSON(200, gin.H{
			"message": "创建用户失败",
		})
	}
	Logger.Info("创建用户成功")
	c.JSON(200, gin.H{
		"message": "创建用户成功",
		"data":    user,
	})

}

func GetUserByUsername(c *gin.Context) {

	var user Model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		Logger.Error("参数绑定失败")
		c.JSON(200, gin.H{
			"message": "参数绑定失败",
		})
	}
	if err := user.GetUserByUsername(); err != nil {
		Logger.Error("获取用户失败")
		c.JSON(200, gin.H{
			"message": "获取用户失败",
		})
	}
	Logger.Info("获取用户成功")
	c.JSON(200, gin.H{
		"message": "获取用户成功",
		"data":    user,
	})

}

func UpdateUser(c *gin.Context) {

	var user Model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		Logger.Error("参数绑定失败")
		c.JSON(200, gin.H{
			"message": "参数绑定失败",
		})
	}
	if err := user.UpdateUser(); err != nil {
		Logger.Error("更新用户失败")
		c.JSON(200, gin.H{
			"message": "更新用户失败",
		})
	}
	Logger.Info("更新用户成功")
	c.JSON(200, gin.H{
		"message": "更新用户成功",
		"data":    user,
	})
}

func GetUserByID(c *gin.Context) {

	var user Model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		Logger.Error("参数绑定失败")
		c.JSON(200, gin.H{
			"message": "参数绑定失败",
		})
	}
	if err := user.GetUserByID(); err != nil {
		Logger.Error("获取用户失败")
		c.JSON(200, gin.H{
			"message": "获取用户失败",
		})
	}
	Logger.Info("获取用户成功")
	c.JSON(200, gin.H{
		"message": "获取用户成功",
		"data":    user,
	})

}

func DeleteUserByID(c *gin.Context) {

	var user Model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		Logger.Error("参数绑定失败")
		c.JSON(200, gin.H{
			"message": "参数绑定失败",
		})
	}
	if err := user.DeleteUserByID(); err != nil {
		Logger.Error("删除用户失败")
		c.JSON(200, gin.H{
			"message": "删除用户失败",
		})
	}
	Logger.Info("删除用户成功")
	c.JSON(200, gin.H{
		"message": "删除用户成功",
		"data":    user,
	})
}

func GetAllUsers(c *gin.Context) {
	var user Model.User
	users, err := user.GetAllUsers()
	if err != nil {
		Logger.Error("获取用户失败")
		c.JSON(200, gin.H{
			"message": "获取用户失败",
		})
	}
	Logger.Info("获取用户成功")
	c.JSON(200, gin.H{
		"message": "获取用户成功",
		"data":    users,
	})
}
