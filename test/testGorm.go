package main

import (
	"chat/model"
	"chat/utils"
)

func main() {

	//生成表
	//utils.DB.AutoMigrate(&model.User{})
	//utils.DB.AutoMigrate(&model.Message{})
	//utils.DB.AutoMigrate(&model.Room{})
	//utils.DB.AutoMigrate(&model.UserRoom{})
	////添加
	user := &model.User{}
	user.Username = "admin1"
	user.Password = "1234"
	user.Email = "admin1@qq.com"
	user.Sex = 1
	//user1.Username = "admin"
	//user1.Password = "12345"
	//user1.Email = "admin@qq.com"
	//user1.Sex = 1
	utils.DB.Create(&user)
	//utils.DB.Create(&user1)
}
