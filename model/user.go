package model

import (
	"time"
)

type User struct {
	Id        uint      `gorm:"comment:'用户id'" json:"id"`
	Username  string    `gorm:"comment:'用户名'" json:"username"`
	Password  string    `gorm:"comment:'密码'" json:"password"`
	Sex       string    `gorm:"comment:'性别'" json:"sex"`
	Email     string    `gorm:"comment:'邮箱'" json:"email"`
	Avatar    string    `gorm:"comment:'头像'" json:"avatar"`
	CreatedAt time.Time `gorm:"comment:'创建时间'" json:"created_at"`
	UpdatedAt time.Time `gorm:"comment:'更新时间'" json:"updated_at"`
}
