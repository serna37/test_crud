package service

import (
	"net/http"
	"test_crud/model"
	"test_crud/sql"

	"github.com/gin-gonic/gin"
)

// ==================
// struct def
// ==================
type chatroom struct {
	sql sql.Chtroom
}
type Chatroom interface {
	Create(c *gin.Context)
	Join(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	RemoveMember(c *gin.Context)
	Delete(c *gin.Context)
}

func NewChatroom() Chatroom {
	return &chatroom{sql.NewChtro()}
}

// ==================
// Imprementation
// ==================

func (chatroom *chatroom) Create(c *gin.Context) {
	var req model.ChatRoomCReq
	usrid := CookieChk(c)
	if usrid == 0 {
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if len(req.RoomName) == 0 || len(req.RoomName) > 60 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	chatroom.sql.Create(req.RoomName, usrid)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "created"})
}

func (chatroom *chatroom) Join(c *gin.Context) {
	var req model.ChatRoomJoinReq
	usrid := CookieChk(c)
	if usrid == 0 {
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.ChatRoomId == 0 || len(req.JoinnersLoginId) == 0 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	resultMsg := chatroom.sql.Join(req.ChatRoomId, req.JoinnersLoginId, usrid)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: resultMsg})
}

func (chatroom *chatroom) Read(c *gin.Context) {
	var req model.ChatRoonRReq
	// XXX cannot carry the cookie from another domein
//	usrid := CookieChk(c)
//	if usrid == 0 {
//		return
//	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	//|| usrid != req.UserId
	if req.UserId == 0 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	rooms := chatroom.sql.Read(req.UserId)
	c.JSON(http.StatusOK, rooms)
}

func (chatroom *chatroom) Update(c *gin.Context) {
	var req model.ChatRoomUReq
	usrid := CookieChk(c)
	if usrid == 0 {
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.RoomId == 0 || len(req.RoomName) == 0 || len(req.RoomName) > 60 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	chatroom.sql.Update(req.RoomId, req.RoomName)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "updated"})
}

func (chatroom *chatroom) RemoveMember(c *gin.Context) {
	var req model.ChatRoomJoinReq
	usrid := CookieChk(c)
	if usrid == 0 {
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.ChatRoomId == 0 || len(req.JoinnersLoginId) == 0 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	resultMsg := chatroom.sql.RemoveMember(req.ChatRoomId, req.JoinnersLoginId)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: resultMsg})
}

func (chatroom *chatroom) Delete(c *gin.Context) {
	var req model.ChatRoomDReq
	usrid := CookieChk(c)
	if usrid == 0 {
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.UserId == 0 || req.UserId != usrid || req.ContentsId == 0 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	chatroom.sql.Delete(req.ContentsId)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "deleted"})
}
