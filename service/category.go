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
type cate struct {
	sql sql.Category
}
type Cate interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewCate() Cate {
	return &cate{sql.NewCategory()}
}

// ==================
// Imprementation
// ==================
func (cate *cate) Create(c *gin.Context) {
	var req model.CategoryCReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.UserId == 0 || len(req.CategoryName) == 0 || len(req.CategoryName) > 60 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	cate.sql.Create(req.UserId, req.CategoryName)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "created"})
}

func (cate *cate) Update(c *gin.Context) {
	var req model.CategoryUReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.CategoryId == 0 || len(req.CategoryName) == 0 || len(req.CategoryName) > 60 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	cate.sql.Update(req.CategoryId, req.CategoryName)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "updated"})
}

func (cate *cate) Delete(c *gin.Context) {
	var req model.CategoryDReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	if req.CategoryId == 0 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		return
	}
	cate.sql.Delete(req.CategoryId)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "deleted"})
}
