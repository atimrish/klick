package ChatController

import (
	"backend/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	helpers.HandleError(err)

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("[websocket] ", err)
			break
		}
		err = conn.WriteMessage(t, msg)
		fmt.Println("[message] ", string(msg))
		helpers.HandleError(err)
	}
}

func CreateChat(c *gin.Context) {
	wsHandler(c.Writer, c.Request)
}
