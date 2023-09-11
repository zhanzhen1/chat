package router

import (
	"chat/middiewares"
	"chat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	ginServer := gin.Default()
	//注册

	//登录
	ginServer.POST("/login", service.Login())
	//是否通过认证
	auth := ginServer.Group("/u", middiewares.AuthCheck())
	//用户详情
	auth.GET("/user/detail", service.UserDetail())
	//查询用户得个人消息
	auth.GET("/user/info", service.UserInfo)
	//发送消息，接收
	auth.GET("/websocket/message", service.WebsocketMessage())
	//聊天记录列表
	auth.GET("/chat/list", service.ChatList)
	//添加好友
	auth.POST("/user/add", service.UserAdd)
	//删除好友
	auth.DELETE("/user/delete", service.UserDelete)
	return ginServer
}
