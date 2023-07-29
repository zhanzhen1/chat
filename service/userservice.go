package service

import (
	"chat/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// Login 用户登录
func Login() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		_, err := dao.Login(username, password)
		if err != nil {
			log.Println("Login()", err)
			return
		}
		context.HTML(http.StatusOK, "index.html", nil)
	}
}
