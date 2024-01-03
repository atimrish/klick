package FriendController

import (
	"backend/controllers/FriendController/types"
	"backend/database/db"
	"backend/database/models/friend"
	"backend/helpers"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	STATUS_WAITING = "WAITING"
	STATUS_ACCEPTED = "ACCEPTED"
	STATUS_DECLINED = "DECLINED"
)

func getInvite(userId, friendId int64) (friend.Friend, bool) {
	query := "SELECT id, user_id, friend_id, status FROM friends WHERE user_id = %d AND  friend_id = %d"
	query = fmt.Sprintf(query, userId, friendId)

	connection := db.PostgresConnection()
	defer connection.Close()

	var tmpFriend friend.Friend

	row := connection.QueryRow(query)
	err := row.Scan(&tmpFriend.Id, &tmpFriend.UserId, &tmpFriend.FriendId, &tmpFriend.Status)

	switch err {
	case sql.ErrNoRows:
		return tmpFriend, false

	default:
		panic(err)
	}

	return  tmpFriend, true
}

func Invite(c *gin.Context)  {
	var form types.InviteForm
	err := c.Bind(&form)
	helpers.HandleError(err)

	accessToken, err := c.Cookie("access_token")
	helpers.HandleError(err)

	payload := helpers.GetPayloadJWT(accessToken)
	form.UserId = payload.UserId

	messages, hasError := form.Validate()

	if hasError {
		c.JSON(422, gin.H{
			"message": "error",
			"errors":  messages,
		})
		return
	}

	_, hasRecord := getInvite(form.UserId, form.FriendId)

	if hasRecord {
		c.JSON(422, gin.H{
			"message": "error",
			"errors":  "вы уже отправили запрос на дружбу с этим пользователем",
		})
		return
	}

	query := "INSERT INTO friends(user_id, friend_id, status) VALUES (%d, %d, '%s')"
	query = fmt.Sprintf(query, form.UserId, form.FriendId, STATUS_WAITING)

	connection := db.PostgresConnection()
	defer connection.Close()

	_, err = connection.Exec(query)
	helpers.HandleError(err)

	c.JSON(201, gin.H{
		"message": "created",
		"errors":  "запрос на дружбу отправлен",
	})
	return
}
