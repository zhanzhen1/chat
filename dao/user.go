package dao

import (
	"chat/model"
	"chat/utils"
	"log"
)

var user *model.User
var userList []*model.User

func Register(username string) (*model.User, error) {
	err := utils.DB.Where("username = ?", username).Find(&user).Error
	if err != nil {
		log.Println("Register()...err", err)
	}
	return user, nil
}

func Login(username string, password string) (*model.User, error) {
	if err := utils.DB.Where("username = ? and password = ?", username, password).
		Find(&user).Error; err != nil {
		log.Println("login...", err)
		return nil, err
	}
	return user, nil
}

// 添加用户
func AddUser(user1 *model.User) error {
	user = &model.User{
		Username: user1.Username,
		Password: user1.Password,
		Sex:      user1.Sex,
		Email:    user1.Email,
	}
	err := utils.DB.Create(&user).Error
	if err != nil {
		log.Println("AddUser()...err", err)
	}
	return nil
}

// 根据id查询用户
func GetUserBasicById(id uint) (*model.User, error) {
	err := utils.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		log.Println("GetUserBasicById()...", err)
		return nil, err
	}
	return user, err
}

// 根据username查询用户消息
func GetUserByUserName(username string) (*model.User, error) {
	sql := "select * from user where  username = ?"
	err := utils.DB.Raw(sql, username).Scan(&user).Error
	if err != nil {
		log.Println("GetUserBasicById()...", err)
		return nil, err
	}
	return user, err
}
