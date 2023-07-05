package sql

import (
	"log"
	"test_crud/model"
	"time"
)

// ==================
// struct def
// ==================
type chtroom struct {
}
type Chtroom interface {
	Create(name string, ownerid int)
	Join(chatroomid int, joinnerloginid string, inviterid int) string
	Read(userid int) []model.TrnChatRoom
	Update(chatroomid int, roomname string)
	RemoveMember(chatroomid int, bandid string) string
	Delete(chatroomid int)
}

func NewChtro() Chtroom {
	return &chtroom{}
}

// ==================
// Imprementation
// ==================
func (chtroom *chtroom) Create(name string, ownerid int) {
	record := model.TrnChatRoom{Name: name, CreaterId: ownerid, DelFlg: false, CreatedAt: time.Now()}
	result := db.Create(&record)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}

func (chtroom *chtroom) Join(chatroomid int, joinnerloginid string, inviterid int) string {
	var targetUser model.MstUsr
	db.Where("usr_login_id = ?", joinnerloginid).First(&targetUser)
	if targetUser.Id == 0 {
		return "invalid value"
	}
	record := model.TrnChatRoomMember{ChatRoomId: chatroomid, JoinerId: targetUser.Id, DelFlg: false, JoinedAt: time.Now(), Inviter: inviterid}
	result := db.Create(&record)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
		return "something went wrong"
	}
	return "joined"
}

func (chtroom *chtroom) Read(userid int) []model.TrnChatRoom {
	var sourcedata []model.TrnChatRoomMember
	db.Where("joiner_id = ? AND del_flg = false", userid).Find(&sourcedata)
	var joinedRoomids []int
	for _, v := range sourcedata {
		joinedRoomids = append(joinedRoomids, v.ChatRoomId)
	}
	var rooms []model.TrnChatRoom
	db.Where("creater_id = ?", userid).Or("id IN ?", joinedRoomids).Find(&rooms)
	return rooms
}

func (chtroom *chtroom) Update(chatroomid int, roomname string) {
	var target model.TrnChatRoom
	db.First(&target, chatroomid)
	target.Name = roomname
	result := db.Save(&target)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}

func (chtroom *chtroom) RemoveMember(chatroomid int, bandid string) string {
	var targetUser model.MstUsr
	db.Where("usr_login_id = ?", bandid).First(&targetUser)
	if targetUser.Id == 0 {
		return "invalid value"
	}
	var target model.TrnChatRoomMember
	db.Where("chat_room_id = ? AND joiner_id = ?", chatroomid, targetUser.Id).First(&target)
	target.DelFlg = true
	result := db.Save(&target)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
		return "something went wrong"
	}
	return "removed"
}

func (chtroom *chtroom) Delete(chatroomid int) {
	var target model.TrnChatRoom
	db.First(&target, chatroomid)
	target.DelFlg = true
	result := db.Save(&target)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}
