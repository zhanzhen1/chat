package model

import "time"

type UserRoom struct {
	UserId    uint      `gorm:"comment:'用户id'"  json:"user_id"`
	RoomId    uint      `gorm:"comment:'房间id'" json:"room_id"`
	RoomType  int       `gorm:"comment:'房间类型'"  json:"room_type"` // 房间 类型 【1-独聊房间 2-群聊房间】
	CreatedAt time.Time `gorm:"comment:'创建时间'"  json:"created_at"`
	UpdatedAt time.Time `gorm:"comment:'更新时间'"  json:"updated_at"`
}
