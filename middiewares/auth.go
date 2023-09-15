package middiewares

import (
	"chat/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 是否通过token
func AuthCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		userClaims, err := helper.AnalyseToken(token)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户认证不通过",
			})
			return
		}
		context.Set("user_claims", userClaims)
		context.Next()
	}
}
