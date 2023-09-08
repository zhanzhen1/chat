package service

import (
	"chat/dao"
	"chat/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// 用户注册
func Register() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {

	}
}

// Login 用户登录
func Login() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		if username == "" || password == "" {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码不能为nil",
			})
			return
		}
		ub, _ := dao.Login(username, password)
		if ub.Id != 0 {
			token, err := helper.GenerateToken(ub.Id, ub.Email)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "系统错误" + err.Error(),
				})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "登录成功",
				"data": gin.H{
					"token": token,
				},
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
		}

	}
}

func UserDetail() gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get("user_claims")
		uc := u.(*helper.UserClaims)
		fmt.Println("id:", uc.Identity)
		userBasic, err := dao.GetUserBasicById(uc.Identity)
		if err != nil {
			log.Panicf("[DB ERROR]:%v\n ", err)
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "数据查询异常",
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": userBasic,
		})

	}
}

func UserInfo(context *gin.Context) {
	username := context.Query("username")
	if username == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户不存在",
		})
		return
	}
	userInfo, err := dao.GetUserByUserName(username)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	if username != userInfo.Username {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名异常",
		})
		return
	}
	uc := context.MustGet("user_claims").(*helper.UserClaims)
	data := &UserInfoResult{
		Username: userInfo.Username,
		Sex:      userInfo.Sex,
		Email:    userInfo.Email,
		IsFriend: false,
	}
	if dao.GetUserIsFriend(userInfo.Id, uc.Identity) {
		data.IsFriend = true
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "数据加载成功",
		"data": data,
	})
}

type UserInfoResult struct {
	Username string `gorm:"comment:'用户名'" json:"username"`
	Sex      int    `gorm:"comment:'性别'" json:"sex"`
	Email    string `gorm:"comment:'邮箱'" json:"email"`
	IsFriend bool   `gorm:"comment:'是否是好友'" json:"IsFriend"` // true 是好友 false 不是
}
