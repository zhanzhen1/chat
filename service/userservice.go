package service

import (
	"chat/dao"
	"chat/helper"
	"chat/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func GetIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// 用户注册
func Register() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		sex := context.PostForm("sex")
		email := context.PostForm("email")
		user, _ := dao.Register(username)
		fmt.Println(user)
		if user.Username == username {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名已存在",
			})
			return
		}
		user1 := &model.User{
			Username: username,
			Password: password,
			Sex:      sex,
			Email:    email,
		}
		err := dao.AddUser(user1)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "添加用户err",
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "注册成功",
		})
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

// 用户详情
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

// 查询用户个人消息
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
	Sex      string `gorm:"comment:'性别'" json:"sex"`
	Email    string `gorm:"comment:'邮箱'" json:"email"`
	IsFriend bool   `gorm:"comment:'是否是好友'" json:"IsFriend"` // true 是好友 false 不是
}

// 添加好友
func UserAdd(context *gin.Context) {
	username := context.PostForm("username")
	if username == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户不存在",
		})
		return
	}
	un, err := dao.GetUserByUserName(username)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	if username != un.Username {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名异常",
		})
		return
	}
	uc := context.MustGet("user_claims").(*helper.UserClaims)
	if dao.GetUserIsFriend(un.Id, uc.Identity) {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "互为好友，无需添加",
		})
		return
	}
	//保存房间记录
	room := &model.Room{
		Id:        helper.GetUUID(),
		UserId:    uc.Identity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = dao.InsertOneRoom(room)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	//保存用户房间的
	ur := &model.UserRoom{
		UserId:    uc.Identity,
		RoomId:    room.Id,
		RoomType:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = dao.InsertOneUserRoom(ur)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	ur = &model.UserRoom{
		UserId:    un.Id,
		RoomId:    room.Id,
		RoomType:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = dao.InsertOneUserRoom(ur)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "添加成功",
	})
}

// 删除好友
func UserDelete(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	uc := context.MustGet("user_claims").(*helper.UserClaims)
	//判断是否为好友
	roomId := dao.GetUserRoomId(id, uc.Identity)
	if roomId == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "不为好友",
		})
		return
	}
	//删除userRoom关系
	err := dao.DeleteUserRoom(roomId)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常",
		})
		return
	}
	err = dao.DeleteRoom(roomId)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "删除成功",
	})
}
