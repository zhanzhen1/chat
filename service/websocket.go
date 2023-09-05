package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

var upgrader = websocket.Upgrader{} // use default options
var ws = make(map[*websocket.Conn]struct{})

func WebsocketMessage(contest *gin.Context) {
	c, err := upgrader.Upgrade(contest.Writer, contest.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	ws[c] = struct{}{}
	//接收到消息
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		//循环拿到coon
		for conn := range ws {
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

//var upgrader = websocket.Upgrader{}
//var wc = make(map[uint]*websocket.Conn)
//
//func WebsocketMessage() (handlerFunc gin.HandlerFunc) {
//	return func(context *gin.Context) {
//		coon, err := upgrader.Upgrade(context.Writer, context.Request, nil)
//		if err != nil {
//			context.JSON(http.StatusOK, gin.H{
//				"code": -1,
//				"msg":  "系统异常:" + err.Error(),
//			})
//			return
//		}
//		defer coon.Close()
//		uc := context.MustGet("user_claims").(*helper.UserClaims)
//		wc[uc.Identity] = coon
//		//读取
//		for {
//			ms := new(define.MessageStruct)
//			err := coon.ReadJSON(ms)
//			if err != nil {
//				log.Printf("Read Error:%v\n", err)
//				return
//			}
//			for _, v := range wc {
//				err := v.WriteMessage(websocket.TextMessage, []byte(ms.Message))
//				if err != nil {
//					log.Printf("write message err:%v\n", err)
//					return
//				}
//			}
//		}
//	}
//}
