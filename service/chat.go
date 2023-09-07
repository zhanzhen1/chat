package service

import (
	"chat/dao"
	"chat/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ChatList(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Query("room_id"), 10, 64)
	roomId := uint(id)
	fmt.Println("room_id:", roomId)
	if roomId == 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "房间号不能为空",
		})
		return
	}
	//判断用户是否属于该房间
	uc := context.MustGet("user_claims").(*helper.UserClaims)
	_, err := dao.GetUserRoomByIdAndRoomId(uc.Identity, roomId)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "房间号非法访问",
		})
		return
	}
	//获取每页数据
	Index, _ := strconv.ParseInt(context.Query("page_index"), 10, 32)
	Size, _ := strconv.ParseInt(context.Query("page_size"), 10, 32)
	pageIndex := int(Index)
	pageSize := int(Size)
	skip := (pageIndex - 1) * pageSize
	//聊天记录查询
	data, err := dao.GetMessageListByRoomId(roomId, pageSize, skip)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "数据加载成功",
		"data": gin.H{
			"list": data,
		},
	})
}
