package model

import (
	"time"
)

type BaseRes struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// ======================================
// Usr
// ======================================
// Create
// Read
// Update
// Delete

// ======================================
// Category
// ======================================
// Create
// Read
type CategoryRRes struct {
	CategoryId   int           `json:"categoryId"`
	CategoryName string        `json:"categoryName"`
	ContentsList []SubContents `json:"contentsList"`
	TagsList     []SubTag      `json:"tagsList"`
}

// Update
// Delete

// ======================================
// Tag
// ======================================
// Create
// Read
type SubTag struct {
	TagId   int    `json:"tagId"`
	TagName string `json:"tagname"`
}

// Update
// Delete

// ======================================
// Contents
// ======================================
// Create
// Read
type SubContents struct {
	Title    string   `json:"title"`
	Contents string   `json:"contents"`
	Tags     []SubTag `json:"tags"`
}

// Update
// Delete

// ======================================
// ChatRoom
// ======================================
// Create
// Read
type ChatRoomRRes struct {
	RoomId   int    `json:"roomId"`
	RoomName string `json:"roomName"`
}

// Update
// Delete

// ======================================
// ChatMsg
// ======================================
// Create
// Read
type ChatMsgRRes struct {
	MsgId  int       `json:"msgId"`
	Msg    string    `json:"msg"`
	FromAt time.Time `json:"fromAt"`
}
