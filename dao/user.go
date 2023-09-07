package dao

import (
	"chat/model"
	"chat/utils"
	"log"
)

var user *model.User
var userList []*model.User

func Login(username string, password string) (*model.User, error) {
	if err := utils.DB.Where("username = ? and password = ?", username, password).
		Find(&user).Error; err != nil {
		log.Println("login...", err)
		return nil, err
	}
	return user, nil
}
func GetUserBasicById(id uint) (*model.User, error) {
	err := utils.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		log.Println("GetUserBasicById()...", err)
		return nil, err
	}
	return user, err
}
