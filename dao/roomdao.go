package dao

import (
	"chat/model"
	"chat/utils"
	"log"
)

var ur *model.UserRoom
var urList []*model.UserRoom

func GetUserRoomByIdAndRoomId(userId, roomId uint) (*model.UserRoom, error) {
	err := utils.DB.Where("user_id = ? and room_id= ?", userId, roomId).Find(&ur).Error
	if err != nil {
		log.Println("GetUserRoomByIdAndRoomId()...", err)
	}
	return ur, err
}

func InsertOneUserRoom(userRoom *model.UserRoom) error {
	err := utils.DB.Create(&userRoom).Error
	if err != nil {
		log.Fatal("Create() err", err)
	}
	return err
}

func InsertOneRoom(room *model.Room) error {
	err := utils.DB.Create(&room).Error
	if err != nil {
		log.Fatal("Create() err", err)
	}
	return err
}

var room *model.Room

func DeleteRoom(id string) error {
	err := utils.DB.Where("id = ? ", id).Delete(&room).Error
	if err != nil {
		log.Println("DeleteRoom()..", err)
	}
	return err
}
