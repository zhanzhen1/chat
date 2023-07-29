package model

type Room struct {
	Id        uint   `gorm:"comment:'用户id'" json:"id"`
	Number    string `gorm:"comment:'房间号'"  json:"number"`
	Name      string `gorm:"comment:'房间名称'" json:"name"`
	Info      string `gorm:"comment:'房间简介'" json:"info"`
	UserId    uint   `gorm:"comment:'用户id'" json:"user_id"`
	CreatedAt int64  ` gorm:"comment:'创建时间'" json:"created_at"`
	UpdatedAt int64  ` gorm:"comment:'更新时间'" json:"updated_at"`
}
