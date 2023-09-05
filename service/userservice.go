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
		ub, err := dao.Login(username, password)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
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
