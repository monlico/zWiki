package wikiWebSocket

import (
	"encoding/json"
	"net/http"
	"sync"
	"zWiki/pkg/logging"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// websocket校验是否挤掉登录
var UserLoginStatusMap = sync.Map{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Data MessageData `json:"data"`
	Type string      `json:"type"`
}

type MessageData struct {
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
	var exitMsg Message = MsgMaker("login_exit")

	for {
		oldUidConn := ReadFromMap(msg.Data.Uid)
		var oldUidConnObj *websocket.Conn
		//旧链接处理
		if oldUidConn != "" {
			oldReadErr := json.Unmarshal([]byte(oldUidConn), &oldUidConnObj)
			if oldReadErr != nil {
				return //连接关闭
			}
			if conn != oldUidConnObj {
				//如果覆盖，顺便将原本的旧链接关闭
				writeToMap(msg.Data.Uid, conn)
				oldUidConnObj.WriteJSON(exitMsg)
				oldUidConnObj.Close()
			}
		} else { //不存在就直接写入
			writeToMap(msg.Data.Uid, conn)
		}
		//校验完连接，开始监听接收到的msg
		conn.ReadJSON(&msg)
		{ //后续处理逻辑

		}
	}
}

func MsgMaker(kind string) Message {
	var msg Message
	switch kind {
	case "exit":
		msg = Message{
			Type: "heartbeat",
			Data: MessageData{},
		}
	case "login_exit":
		msg = Message{
			Type: "close",
			Data: MessageData{},
		}
	default:
		msg = Message{
			Type: "default",
			Data: MessageData{},
		}
	}
	return msg
}

func writeToMap(key, value interface{}) {
	UserLoginStatusMap.Store(key, value)
}

func ReadFromMap(key interface{}) string {
	value, _ := UserLoginStatusMap.Load(key)
	if value != nil {
		return value.(string)
	}
	return ""
}
