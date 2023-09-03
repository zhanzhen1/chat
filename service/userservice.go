package service

import (
	"chat/dao"
	"github.com/gin-gonic/gin"
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
				"msg":  "用户名或密码不能为空",
			})
			return
		}
		_, err := dao.Login(username, password)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "success",
		})
	}
}
