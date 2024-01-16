package ChatController

import (
	"backend/controllers/ChatController/types"
	"backend/database/models/chat"
	"backend/database/models/message"
	"backend/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func GetChatsByUserId(c *gin.Context) {
	userId := c.Param("user_id")

	userIdInt, err := strconv.Atoi(userId)
	helpers.HandleError(err)

	chats, err := chat.GetByUserId(int64(userIdInt))
	if err != nil {
		helpers.ErrorResponse(c, 422, err)
	}

	c.JSON(200, gin.H{
		"data": *chats,
	})
	return
}

func GetChatById(c *gin.Context) {
	chatId := c.Param("chat_id")

	returnedChat, err := chat.GetById(chatId)
	if err != nil {
		helpers.ErrorResponse(c, 422, err)
	}

	c.JSON(200, gin.H{
		"data": *returnedChat,
	})
	return
}

func CreateChat(c *gin.Context) {
	var form types.CreateChatForm
	err := c.Bind(&form)
	helpers.HandleError(err)

	usersArray := form.GetUsersArray()

	var newChat chat.Chat

	newChat.Type = "PERSONAL"
	newChat.Users = *usersArray
	newChat.Messages = []message.Message{}
	go func(c *gin.Context) {
		newChat.Insert()
		c.JSON(200, gin.H{
			"data": gin.H{
				"chatId": newChat.Id,
			},
		})
		return
	}(c)
}

func PushMessage(c *gin.Context) {
	chatId := c.Param("chat_id")

	var form types.AddMessageForm
	err := c.Bind(&form)
	helpers.HandleError(err)

	fmt.Println("[form] ", form)
	fmt.Println("[ctx] ", c.Request)

	var newMessage message.Message
	newMessage.UserId = form.UserId
	newMessage.Text = form.Text

	objId, err := primitive.ObjectIDFromHex(chatId)
	helpers.HandleError(err)

	go func(c *gin.Context) {
		newMessage.PushMessage(objId)
		c.JSON(201, gin.H{
			"message": "сообщение отправлено",
		})
		return
	}(c)
}

func UpdateMessage(c *gin.Context) {
	chatId := c.Param("chat_id")
	messageId := c.Param("message_id")

	var form types.UpdateMessageForm
	err := c.Bind(&form)
	helpers.HandleError(err)

	objId, err := primitive.ObjectIDFromHex(chatId)
	helpers.HandleError(err)

	go func(c *gin.Context) {
		message.UpdateMessage(objId, messageId, form.Text)
		c.JSON(200, gin.H{
			"message": "сообщение обновлено",
		})
		return
	}(c)
}

func DeleteMessage(c *gin.Context) {
	chatId := c.Param("chat_id")
	messageId := c.Param("message_id")

	objId, err := primitive.ObjectIDFromHex(chatId)
	helpers.HandleError(err)

	go func(c *gin.Context) {
		message.DeleteMessage(objId, messageId)
		c.JSON(200, gin.H{
			"message": "сообщение удалено",
		})
		return
	}(c)
}
