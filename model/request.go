package model

import (
	"time"
)

// ======================================
// Usr
// ======================================
// Create
type SignupReq struct {
	Name     string `json:"name"`
	Loginid  string `json:"loginid"`
	Password string `json:"password"`
}

// Read
type SigninReq struct {
	Loginid  string `json:"loginid"`
	Password string `json:"password"`
}

// Update
type UserEditReq struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Loginid  string `json:"loginid"`
	Password string `json:"password"`
}

// Delete
type SignoutReq struct {
	Loginid string `json:"loginid"`
}

// ==========================-
// common
type GetRequest struct {
	Id int `json:"id"`
}

// ======================================
// Category
// ======================================
// Create
type CategoryCReq struct {
	UserId       int    `json:"userId"`
	CategoryName string `json:"categoryName"`
}

// Read = response
// Update
type CategoryUReq struct {
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

// Delete
type CategoryDReq struct {
	CategoryId int `json:"categoryId"`
}

// ======================================
// Tag
// ======================================
// Create
type TagCReq struct {
	UserId       int    `json:"userId"`
	CategoryId int    `json:"categoryId"`
	TagName    string `json:"tagname"`
}

// Read = response
// Update
type TagUReq struct {
	TagId   int    `json:"tagId"`
	TagName string `json:"tagname"`
}

// Delete
type TagDReq struct {
	TagId int `json:"tagId"`
}

// ======================================
// Contents
// ======================================
// Create
type ContentsCReq struct {
	UserId   int    `json:"userId"`
	CategoryId int    `json:"categoryId"`
	Title      string `json:"title"`
	Contents   string `json:"contents"`
}

// Read = response
// Update
type ContentsUReq struct {
	ContentsId int    `json:"contentsId"`
	Title      string `json:"title"`
	Contents   string `json:"contents"`
}

// Delete
type ContentsDReq struct {
	ContentsId int `json:"contentsId"`
}

// ======================================
// ChatRoom
// ======================================
// Create
type ChatRoomCReq struct {
	RoomName  string    `json:"roomName"`
	CreaterId int       `json:"createrId"`
	CreatedAt time.Time `json:"createdAt"`
}

// Read
type ChatRoonRReq struct {
	UserId int `json:"userId"`
}

// Update
type ChatRoomUReq struct {
	UserId   int    `json:"userId"`
	RoomId   int    `json:"roomId"`
	RoomName string `json:"roomName"`
}

type ChatRoomJoinReq struct {
	ChatRoomId int `json:"roomId"`
	JoinnersLoginId string `json:"joinnerloginid"`
}

// Delete
type ChatRoomDReq struct {
	UserId     int `json:"userId"`
	ContentsId int `json:"contentsId"`
}

// ======================================
// ChatMsg
// ======================================
// Create
type ChatMsgCReq struct {
	UserId string       `json:"userId"`
	Msg    string    `json:"msg"`
}

// Read
type ChatMsgRReq struct {
	RoomId int `json:"roomId"`
}
