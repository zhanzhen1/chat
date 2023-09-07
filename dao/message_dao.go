package dao

import (
	"chat/model"
	"chat/utils"
	"log"
)

var message *model.Message
var messageList []*model.Message

func InsertMessage(message *model.Message) error {
	message = &model.Message{
		UserId:    message.UserId,
		RoomId:    message.RoomId,
		Data:      message.Data,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
	err := utils.DB.Create(&message).Error
	if err != nil {
		utils.DB.Rollback()
		log.Fatal("AddBook() err:", err)
	}
	return nil
}

// 查询聊天记录
func GetMessageListByRoomId(RoomId uint, limit, skip int) ([]*model.Message, error) {
	data := make([]*model.Message, 0)
	err := utils.DB.Where("room_id = ?", RoomId).Find(&messageList).Error
	if err != nil {
		log.Println("GetMessageListByRoomId()...", err)
		return nil, err
	}

	utils.DB.Limit(limit).Offset(skip).Find(&messageList)
	//将数据append到data
	data = append(data, messageList...)

	return data, err
}
