package main

import (
	"chat/dao"
	"fmt"
	"testing"
)

func TestUserLogin(t *testing.T) {
	user, err := dao.Login("admin", "12345")
	if user.Id == 0 {
		fmt.Println("err", err)
		return
	}
	fmt.Println("user:", user)
}

func TestUserInfo(t *testing.T) {
	userInfo, err := dao.GetUserByUserName("zz")
	if err != nil {
		fmt.Println("GetUserByUserName()...", err)
		return
	}
	fmt.Println("user:", userInfo)

}
