package middiewares

import (
	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.GetHeader("token")
	}
}
