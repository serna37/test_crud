package sql

import (
	"log"
	"test_crud/model"
	"time"
)

// ==================
// struct def
// ==================
type chatmsg struct {
}
type Chatmsg interface {
	SelectFromRoomId(roomid int) []model.TrnChatMsg
	InsertMsg(roomid int, fromid int, msg string)
}

func NewChatmsg() Chatmsg {
	return &chatmsg{}
}

// ==================
// Imprementation
// ==================
func (chatmsg *chatmsg) SelectFromRoomId(roomid int) []model.TrnChatMsg {
	var rows []model.TrnChatMsg
	db.Where("chat_room_id = ?", roomid).Find(&rows)
	return rows
}

func (chatmsg *chatmsg) InsertMsg(roomid int, fromid int, msg string) {
	record := model.TrnChatMsg{ChatRoomId: roomid, FromId: fromid, Msg: msg, FromAt: time.Now()}
	result := db.Create(&record)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}
