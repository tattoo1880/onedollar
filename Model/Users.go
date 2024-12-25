package Model

import (
	"gorm.io/gorm"
	. "onedollarback/Config"
)

type User struct {
	gorm.Model
	ID        uint    `json:"id" gorm:"primary_key"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	LoginTime *string `json:"login_time"`
}

// CreateUser 创建用户
func (u *User) CreateUser() error {
	return MysqlDB.Create(u).Error
}

// GetUserByUsername 根据用户名和密码获取用户
func (u *User) GetUserByUsername() error {
	return MysqlDB.Where("username = ? AND password = ?", u.Username, u.Password).First(u).Error
}

// UpdateUser 更新用户信息
func (u *User) UpdateUser() error {
	return MysqlDB.Model(u).Updates(u).Error
}

// GetUserByID 根据ID获取用户
func (u *User) GetUserByID() error {
	return MysqlDB.Where("id = ?", u.ID).First(u).Error
}

// DeleteUserByID 根据ID删除用户
func (u *User) DeleteUserByID() error {
	return MysqlDB.Where("id = ?", u.ID).Delete(u).Error
}

// GetAllUsers 获取所有用户
func (u *User) GetAllUsers() ([]User, error) {
	var users []User
	err := MysqlDB.Find(&users).Error
	return users, err
}
