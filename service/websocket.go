package service

import (
	"chat/dao"
	"chat/define"
	"chat/helper"
	"chat/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}
var wc = make(map[uint]*websocket.Conn)

func WebsocketMessage() (handlerFunc gin.HandlerFunc) {
	return func(context *gin.Context) {
		coon, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "系统异常:" + err.Error(),
			})
			return
		}
		defer coon.Close()
		uc := context.MustGet("user_claims").(*helper.UserClaims)
		wc[uc.Identity] = coon
		//读取
		for {
			ms := new(define.MessageStruct)
			err := coon.ReadJSON(ms)
			if err != nil {
				log.Printf("Read Error:%v\n", err)
				return
			}
			//判断是否属于消息体的房间
			_, err = dao.GetUserRoomByIdAndRoomId(uc.Identity, ms.RoomId)
			if err != nil {
				log.Printf("userID:%v RoomId:%v\n", uc.Identity, ms.RoomId)
				return
			}
			fmt.Println("uc.id:", uc.Identity)
			fmt.Println("ms.RoomId:", ms.RoomId)
			//保存消息
			msg := &model.Message{
				UserId:    uc.Identity,
				RoomId:    ms.RoomId,
				Data:      ms.Message,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			err = dao.InsertMessage(msg)
			if err != nil {
				log.Println("InsertMessage...", err)
				return
			}
			for _, v := range wc {
				err := v.WriteMessage(websocket.TextMessage, []byte(ms.Message))
				if err != nil {
					log.Printf("write message err:%v\n", err)
					return
				}
			}
		}
	}
}
