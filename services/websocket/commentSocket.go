package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"zWiki/controller/wikiWebSocket"
)

// 评论时的socket处理
func AddCommentSocket(receiver, send string) {
	webConn := wikiWebSocket.ReadFromMap(receiver)

	var webConnObj *websocket.Conn

	json.Unmarshal([]byte(webConn), &webConnObj)

	var msg = wikiWebSocket.MsgMaker("addComment")

	msg.Data.Token = send

	webConnObj.WriteJSON(msg)
}
