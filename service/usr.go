package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"test_crud/model"
	"test_crud/sql"
)

// ==================
// struct def
// ==================
type usr struct {
	sql sql.Usr
}
type Usr interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
}

func NewUsr() Usr {
	return &usr{sql.NewUsr()}
}

// ==================
// Imprementation
// ==================

func (usr *usr) Signup(c *gin.Context) {
	log.Printf("Signup start")

	var req model.SignupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		log.Printf("Signup end")
		return
	}

	userChk := usr.sql.Check(req.Loginid)
	if !userChk {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "duplicated id"})
		log.Printf("Signup end")
		return
	}

	// TODO ランダムにならない:
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 乱数を生成
	b := make([]byte, 128)
	if _, err := rand.Read(b); err != nil {
		log.Printf("Unexpected error")
	}
	// letters からランダムに取り出して文字列を生成
	var authtoken string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		authtoken += string(letters[int(v)%len(letters)])
	}

	usr.sql.Create(req.Name, req.Loginid, req.Password, authtoken)

	cookie := authtoken
	option := gin.H{"domain": "serna37.github.io", "path": "/", "sameSite": "None"}
	c.JSON(http.StatusOK, gin.H{"status": 0, "cookie": cookie, "option": option})
	log.Printf("Signup end")
}

func (usr *usr) Signin(c *gin.Context) {
	log.Printf("Signin start")

	var req model.SigninReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		log.Printf("Signin end")
		return
	}

	userdata := usr.sql.Read(req.Loginid, req.Password)
	if userdata.Id == -1 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "no match"})
		log.Printf("Signin end")
		return
	}

	usr.sql.Update(userdata.Id, "", "", "")
	cookie := userdata.AuthToken
	option := gin.H{"domain": "serna37.github.io", "path": "/", "sameSite": "None"}
	c.JSON(http.StatusOK, gin.H{"status": 0, "cookie": cookie, "option": option})
	log.Printf("Signin end")
}
