package router

import (
	"chat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	ginServer := gin.Default()
	ginServer.LoadHTMLGlob("./view/*.html")
	ginServer.Static("/static", "./static")
	ginServer.POST("/login", service.Login())
	return ginServer
}
