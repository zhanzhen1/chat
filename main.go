package main

import (
	"chat/router"
	"chat/utils"
)

func main() {
	utils.InitRedis()
	ginServer := router.Router()
	ginServer.Run(":8080")
}
