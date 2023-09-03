package router

import (
	"chat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	ginServer := gin.Default()
	ginServer.POST("/login", service.Login())
	return ginServer
}
