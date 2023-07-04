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
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetCateTag(c *gin.Context)
	GetAllData(c *gin.Context)
}

func NewUsr() Usr {
	return &usr{sql.NewUsr()}
}

// ==================
// Imprementation
// ==================

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func (usr *usr) Signup(c *gin.Context) {
	log.Printf("Signup start")

	var req model.SignupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		log.Printf("Signup end")
		return
	}

	// validation
	if len(req.Loginid) == 0 || len(req.Password) == 0 || len(req.Name) == 0 || len(req.Loginid) > 63 || len(req.Name) > 63 || len(req.Password) > 1000 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		log.Printf("Signup end")
		return
	}
	userChk := usr.sql.Check(req.Loginid)
	if !userChk {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "duplicated id"})
		log.Printf("Signup end")
		return
	}

	authtoken := RandomString(10)
	usr.sql.Create(req.Name, req.Loginid, req.Password, authtoken)

	option := gin.H{"domain": "serna37.github.io", "path": "/", "sameSite": "None"}
	c.JSON(http.StatusOK, gin.H{"status": 0, "cookie": authtoken, "option": option})
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
	if userdata.Id == 0 {
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

func (usr *usr) Update(c *gin.Context) {
	log.Printf("userinfo edit start")

	var req model.UserEditReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		log.Printf("userinfo edit end")
		return
	}

	// validation
	if len(req.Loginid) == 0 || len(req.Password) == 0 || len(req.Name) == 0 || len(req.Loginid) > 63 || len(req.Name) > 63 || len(req.Password) > 1000 {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "invalid value"})
		log.Printf("userinfo edit end")
		return
	}
	userChk := usr.sql.Check(req.Loginid)
	if !userChk {
		c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "duplicated id"})
		log.Printf("userinfo edit end")
		return
	}
	usr.sql.Update(req.Id, req.Name, req.Loginid, req.Password)
	c.JSON(http.StatusOK, model.BaseRes{Status: 0, Message: "updated"})
	log.Printf("userinfo edit end")
}

func (usr *usr) Delete(c *gin.Context) {
	// no function
	c.JSON(http.StatusOK, model.BaseRes{Status: 1, Message: "this function doesnot supported"})
}

func (usr *usr) GetCateTag(c *gin.Context) {
	var req model.GetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	response := usr.sql.GetCateTag(req.Id)
	c.JSON(http.StatusOK, response)
}

func (usr *usr) GetAllData(c *gin.Context) {
	var req model.GetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.BaseRes{Status: 1, Message: err.Error()})
		return
	}
	response := usr.sql.GetAllData(req.Id)
	c.JSON(http.StatusOK, response)
}
