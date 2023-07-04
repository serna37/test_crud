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
type tag struct {
	sql sql.Tag
}
type Tag interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewTag() Tag {
	return &tag{sql.NewTag()}
}

// ==================
// Imprementation
// ==================
func (tag *tag) Create(c *gin.Context) {
	var req model.TagCReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.UserId == 0 || req.CategoryId == 0 || len(req.TagName) == 0 || len(req.TagName) > 60 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
	}
	tag.sql.Create(req.UserId, req.TagName, req.CategoryId)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "created"})
}

func (tag *tag) Update(c *gin.Context) {
	var req model.TagUReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.TagId == 0 || len(req.TagName) == 0 || len(req.TagName) > 60 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
	}
	tag.sql.Update(req.TagId, req.TagName)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "updated"})
}

func (tag *tag) Delete(c *gin.Context) {
	var req model.TagDReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.TagId == 0 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
	}
	tag.sql.Delete(req.TagId)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "deleted"})
}
