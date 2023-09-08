package dao

import (
	"chat/model"
	"chat/utils"
	"log"
)

var userRoom *model.UserRoom
var userRoonList []*model.UserRoom

// 根据room_type判断是否为好友
func GetUserIsFriend(userId1, userId2 uint) bool {
	//查询userId1的房间
	err := utils.DB.Where("user_id = ? and room_type = ?", userId1, 1).Find(&userRoom).Error
	if err != nil {
		log.Println("GetUserIsFriend()..userId1 err	", err)
		return false
	}
	if userRoom.RoomType == 0 {
		return false
	}
	//获取关联 userId2的房间
	err = utils.DB.Where("user_id = ? and room_type = ?", userId2, userRoom.RoomType).Find(&userRoom).Error
	if err != nil {
		log.Println("GetUserIsFriend()..userId2 err	", err)
		return false
	}
	return true
}
