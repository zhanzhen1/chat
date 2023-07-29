package model

import "time"

// Message 消息
type Message struct {
	UserId    uint      `gorm:"comment:'用户id'" json:"user_identity"`
	RoomId    uint      `gorm:"comment:'房间id'" json:"room_identity"`
	Data      string    `gorm:"comment:'消息数据'" json:"data"`
	CreatedAt time.Time `gorm:"comment:'创建时间'" json:"created_at"`
	UpdatedAt time.Time `gorm:"comment:'更新时间'" json:"updated_at"`
}
