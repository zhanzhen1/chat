package define

import "os"

var MailPassword = os.Getenv("MailPassword")

type MessageStruct struct {
	Message string `json:"message"`
	RoomId  string `json:"roomId"`
}
