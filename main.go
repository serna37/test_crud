package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"test_crud/service"
	"test_crud/sql"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("hello")

	// create DB connection
	sql.Conn()

	r := gin.Default()

	// create CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
			"https://serna37.github.io",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS", // for preflight request
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,           // need cookie
		MaxAge:           24 * time.Hour, // preflight request's result chache term
	}))

	// regist API endpoints
	rg := r.Group("/mk6v2")

	// usr
	usr := service.NewUsr()
	rg1 := rg.Group("/usr")
	rg1.POST("/signup", usr.Signup)
	rg1.POST("/signin", usr.Signin)
	rg1.POST("/update", usr.Update)
	rg1.POST("/delete", usr.Delete)
	rg1.POST("/getcatetag", usr.GetCateTag)
	rg1.POST("/getalldata", usr.GetAllData)

	// category
	cate := service.NewCate()
	rg2 := rg.Group("/category")
	rg2.POST("/create", cate.Create)
	rg2.POST("/update", cate.Update)
	rg2.POST("/delete", cate.Delete)

	// tag
	tag := service.NewTag()
	rg3 := rg.Group("/tag")
	rg3.POST("/create", tag.Create)
	rg3.POST("/update", tag.Update)
	rg3.POST("/delete", tag.Delete)

	// contents
	contnt := service.NewContnt()
	rg4 := rg.Group("/contents")
	rg4.POST("/create", contnt.Create)
	rg4.POST("/update", contnt.Update)
	rg4.POST("/delete", contnt.Delete)

	// chatroom
	chtrm := service.NewChatroom()
	rg5 := rg.Group("/chatroom")
	rg5.POST("/create", chtrm.Create)
	rg5.POST("/join", chtrm.Join)
	rg5.POST("/read", chtrm.Read)
	rg5.POST("/update", chtrm.Update)
	rg5.POST("/remove", chtrm.RemoveMember)
	rg5.POST("/delete", chtrm.Delete)

	// chat
	service.WSIni()
	rg.GET("/msg/:roomid", service.MsgSync)

	// weather
	rg.GET("/weather", test)

	// ping
	rg.POST("/ping", ping)

	r.Run(":8181")
}

// for test
func ping(c *gin.Context) {
	//	uu := sql.NewUsr()
	//	uu.GetAllData(1)
	c.JSON(http.StatusOK, gin.H{"status": 0})
}

func test(c *gin.Context) {

	openweatherendpoint := "http://api.openweathermap.org/data/2.5/forecast?q=Tokyo,JP&appid=55426606e82455879b287daa26fc9204&lang=ja&units=metric"

	res, err := http.Get(openweatherendpoint)
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	// []byte -> string
	c.JSON(http.StatusOK, gin.H{"jsondata": string(body)})
}
