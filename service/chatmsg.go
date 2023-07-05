package service

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
	"test_crud/model"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"test_crud/sql"
)

var m *melody.Melody

func WSIni() {
	m = melody.New()
	lock := new(sync.Mutex)
	chatmsg := sql.NewChatmsg()

	// on connected
	m.HandleConnect(func(s *melody.Session) {
		log.Printf("#### websocket connection open. [session: %#v]\n", s)
		lock.Lock()

		// get roomid
		reqUris := strings.Split(s.Request.URL.String(), "/")
		roomId := reqUris[len(reqUris)-1]
		i, _ := strconv.Atoi(roomId)

		// show all msg
		msgs := chatmsg.SelectFromRoomId(i)
		jsonstring, err := json.Marshal(msgs)
		if err != nil {
			log.Fatal(err.Error())
		}
		s.Write(jsonstring)

		lock.Unlock()
	})

	// on disconnected
	m.HandleDisconnect(func(s *melody.Session) {
		message := "#### websocket disconnected."
		s.Write([]byte(message))
		log.Printf("#### websocket connection close. [session: %#v]\n", s)
	})

	// on message
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		lock.Lock()

		// get roomid
		reqUris := strings.Split(s.Request.URL.String(), "/")
		roomId := reqUris[len(reqUris)-1]
		i, _ := strconv.Atoi(roomId)

		// insert reqeust
		var req model.ChatMsgCReq
		if e := json.Unmarshal(msg, &req); e != nil {
			log.Fatal(e.Error())
		}
		userid, _ := strconv.Atoi(req.UserId)
		chatmsg.InsertMsg(i, userid, req.Msg)

		// send message for same room
		sessions, err := m.Sessions()
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, v := range sessions {
			uri := strings.Split(v.Request.URL.String(), "/")
			joindroomId := uri[len(uri)-1]
			// broadcast only same room (including me)
			if roomId == joindroomId {
				v.Write(msg)
			}
		}

		message := string(msg)
		log.Printf("[accept message]:%s", message)
		//m.BroadcastOthers(msg, s)
		//m.Broadcast(msg)
		lock.Unlock()
	})
}

func MsgSync(c *gin.Context) {
	m.HandleRequest(c.Writer, c.Request)
}
