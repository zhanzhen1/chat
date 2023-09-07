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
