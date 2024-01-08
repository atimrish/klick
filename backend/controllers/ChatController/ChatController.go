package ChatController

import (
	"backend/database/models/chat"
	"backend/database/models/message"
	"backend/helpers"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
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

		PushMessage(msg)
		fmt.Println("[message] ", t, " ", string(msg))
		helpers.HandleError(err)
	}
}

func CreateChat(c *gin.Context) {
	wsHandler(c.Writer, c.Request)
}

func PushMessage(income []byte) {

	type chatId struct {
		ChatId   string `json:"chat_id"`
		Receiver int64  `json:"receiver"`
	}
	var chatInfo chatId

	var m message.Message
	err := json.Unmarshal(income, &m)
	helpers.HandleError(err)

	err = json.Unmarshal(income, &chatInfo)
	helpers.HandleError(err)

	if chatInfo.ChatId == "" {
		users := []int64{m.UserId, chatInfo.Receiver}

		m.Id = uuid.New().String()
		m.CreatedTime = strconv.Itoa(int(time.Now().Unix()))

		messages := []message.Message{m}

		newChat := chat.Chat{
			Users:    users,
			Title:    "",
			Messages: messages,
			Photo:    "",
			Type:     "PRIVATE",
		}

		go newChat.Insert()
		return
	}

	objId, err := primitive.ObjectIDFromHex(chatInfo.ChatId)
	helpers.HandleError(err)

	fmt.Println("[obj id] ", objId)
	m.PushMessage(objId)
	return
}
