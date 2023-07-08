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
type contn struct {
	sql sql.Contents
}
type Contnt interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewContnt() Contnt {
	return &contn{sql.NewContents()}
}

// ==================
// Imprementation
// ==================
func (contn *contn) Create(c *gin.Context) {
//	usrid := CookieChk(c)
	// XXX cannot carry the cookie from another domein
//	if usrid == 0 {
//		return
//	}
	var req model.ContentsCReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.UserId == 0 || req.CategoryId == 0 || len(req.Title) == 0 || len(req.Title) > 60 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	contn.sql.Create(req.UserId, req.Contents, req.CategoryId, req.Title)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "created"})
}

func (contn *contn) Update(c *gin.Context) {
	// XXX cannot carry the cookie from another domein
//	usrid := CookieChk(c)
//	if usrid == 0 {
//		return
//	}
	var req model.ContentsUReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.ContentsId == 0 || len(req.Title) == 0 || len(req.Title) > 60 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	contn.sql.Update(req.Contents, req.Title, req.ContentsId)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "updated"})
}

func (contn *contn) Delete(c *gin.Context) {
	usrid := CookieChk(c)
	if usrid == 0 {
		return
	}
	var req model.ContentsDReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.ContentsId == 0 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	contn.sql.Delete(req.ContentsId)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "deleted"})
}
