package wikiWebSocket

import (
	"net/http"
	"sync"
	"zWiki/pkg/logging"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//websocket校验是否挤掉登录
var UserLoginStatusMap = sync.Map{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Uid    uint   `json:"uid"`
	Action string `json:"action"`
	Token  string `json:"token"`
}

func HandleLoginConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logging.Error(err)
		return
	}
	defer conn.Close()

	var msg Message
	for {
		err := conn.ReadJSON(&msg)
		if err != nil {
			logging.Error(err)
			return
		}
		writeToMap(msg.Uid, msg.Token)
		for {
			if readFromMap(msg.Uid) != msg.Token {
				msg.Action = "exit"
				err := conn.WriteJSON(msg)
				if err != nil {
					logging.Error(err)
					return
				}
				return
			}
		}
	}
}

func writeToMap(key, value interface{}) {
	UserLoginStatusMap.Store(key, value)
}

func readFromMap(key interface{}) string {
	value, _ := UserLoginStatusMap.Load(key)
	if value != nil {
		return value.(string)
	}
	return ""
}
