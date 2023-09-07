package router

import (
	"chat/middiewares"
	"chat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	ginServer := gin.Default()
	//登录
	ginServer.POST("/login", service.Login())
	auth := ginServer.Group("/u", middiewares.AuthCheck())
	//用户详情
	auth.GET("/user/detail", service.UserDetail())
	//发送消息，接收
	auth.GET("/websocket/message", service.WebsocketMessage())
	//聊天记录列表
	auth.GET("/chat/list", service.ChatList)
	return ginServer
}
