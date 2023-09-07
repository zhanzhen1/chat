package define

import "os"

var MailPassword = os.Getenv("MailPassword")

type MessageStruct struct {
	Message string `json:"message"`
	RoomId  uint   `json:"room_id"`
}
