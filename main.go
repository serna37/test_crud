package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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
			"http://mk6.neras-sta.com",
			"https://mk6.neras-sta.com",
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
	rg1 := rg.Group("/usr")
	usr := service.NewUsr()
	rg1.POST("/signup", usr.Signup)
	rg1.POST("/signin", usr.Signin)

	r.Run(":8181")
}

// TODO call API sample
func test(c *gin.Context) {
	var m = [...]string{
		"AAAAAAAAA",
		"AAAAAAAAA",
		"AAAAAAAAA",
		"AAAAAAAAA",
		"AAAAAAAAA",
		"AAAAAAAAA",
		"AAAAAAAAA",
	}
	fmt.Printf(strings.Join(m[:], ","))

	url := "https://neras-sta.com/mk6/getdata"
	authHeaderName := "x-cdata-authtoken"
	authHeaderValue := "7y3E6q4b6V1v9f0D2m9j"

	req, _ := http.NewRequest(http.MethodPost, url, nil)
	//req.Host = "neras-sta.com"
	req.Header.Set(authHeaderName, authHeaderValue)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Request:", err)
		return
	}

	type Contents struct {
		Id        int      `json:"id"`
		Category  string   `json:"category"`
		Thumbnail string   `json:"thumbnail"`
		Title     string   `json:"title"`
		DataUrl   []string `json:"dataUrl"`
		ViewCnt   int      `json:"viewCnt"`
		LikeCnt   int      `json:"likeCnt"`
		Tags      []string `json:"tags"`
	}

	type Res struct {
		Contents []Contents
		tags     []string
	}

	defer resp.Body.Close()
	var response Res
	json.NewDecoder(resp.Body).Decode(&response)
	log.Printf(response.Contents[0].DataUrl[0])
}
