package FriendController

import (
	"backend/controllers/FriendController/types"
	"backend/database/db"
	"backend/database/models/friend"
	"backend/helpers"
	"database/sql"
	"errors"
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

	if errors.Is(err, sql.ErrNoRows) {
		return tmpFriend, false
	}
	helpers.HandleError(err)

	return  tmpFriend, true
}

func getInviteById(id int64) (friend.Friend, bool) {
	query := "SELECT id, user_id, friend_id, status FROM friends WHERE id = %d"
	query = fmt.Sprintf(query, id)

	connection := db.PostgresConnection()
	defer connection.Close()

	var tmpFriend friend.Friend

	row := connection.QueryRow(query)
	err := row.Scan(&tmpFriend.Id, &tmpFriend.UserId, &tmpFriend.FriendId, &tmpFriend.Status)

	if errors.Is(err, sql.ErrNoRows) {
		return tmpFriend, false
	}
	helpers.HandleError(err)

	return  tmpFriend, true
}

func Invite(c *gin.Context)  {
	var form types.InviteForm
	err := c.Bind(&form)
	helpers.HandleError(err)
	form.Status = STATUS_WAITING

	accessToken, err := c.Cookie("access_token")
	helpers.HandleError(err)

	payload, err := helpers.GetPayloadJWT(accessToken)
	if err != nil {
		helpers.TokenExpiredResponse(c)
		return
	}

	form.UserId = payload.UserId

	messages, hasError := form.Validate()
	fmt.Println("has error: ", hasError)

	if hasError {
		c.JSON(422, gin.H{
			"message": "error",
			"errors":  messages,
		})
		return
	}

	_, hasRecord := getInvite(form.UserId, form.FriendId)
	fmt.Println("has record: ", hasRecord)

	if hasRecord {
		c.JSON(422, gin.H{
			"message": "error",
			"errors":  "вы уже отправили запрос на дружбу с этим пользователем",
		})
		return
	}

	query := "INSERT INTO friends(user_id, friend_id, status) VALUES (%d, %d, '%s')"
	query = fmt.Sprintf(query, form.UserId, form.FriendId, form.Status)

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

func Accept(c *gin.Context)  {
	var form types.AcceptForm
	err := c.Bind(&form)
	helpers.HandleError(err)

	accessToken, err := c.Cookie("access_token")
	helpers.HandleError(err)

	payload, err := helpers.GetPayloadJWT(accessToken)
	if err != nil {
		helpers.TokenExpiredResponse(c)
		return
	}

	UserId := payload.UserId

	invite, hasRecord := getInviteById(form.InviteId)

	if !hasRecord {
		c.JSON(422, gin.H{
			"message": "error",
			"errors":  "вы еще не отправляли запроса дружбы данному пользователю",
		})
		return
	}

	if invite.FriendId != UserId {
		c.JSON(403, gin.H{
			"message": "error",
			"errors":  "вы не тот пользователь",
		})
		return
	}

	invite.Status = STATUS_ACCEPTED
	go invite.Update()

	c.JSON(200, gin.H{
		"message": "ok",
		"errors":  "вы подтвердили что вы друзья",
	})

	return
}

func Decline(c *gin.Context) {
	var form types.DeclineForm
	err := c.Bind(&form)
	helpers.HandleError(err)

	accessToken, err := c.Cookie("access_token")
	helpers.HandleError(err)

	payload, err := helpers.GetPayloadJWT(accessToken)
	if err != nil {
		helpers.TokenExpiredResponse(c)
		return
	}

	UserId := payload.UserId

	invite, hasRecord := getInviteById(form.InviteId)

	if !hasRecord {
		c.JSON(422, gin.H{
			"message": "error",
			"errors":  "вы еще не отправляли запроса дружбы данному пользователю",
		})
		return
	}

	if invite.FriendId != UserId {
		c.JSON(403, gin.H{
			"message": "error",
			"errors":  "вы не тот пользователь",
		})
		return
	}

	invite.Status = STATUS_DECLINED
	go invite.Update()

	c.JSON(200, gin.H{
		"message": "ok",
		"errors":  "вы отклонили заявку в друзья",
	})

	return
}

func UserFriends(c *gin.Context) {
	userId := c.Param("user_id")

	connection := db.PostgresConnection()

	defer func() {
		err := connection.Close()
		helpers.HandleError(err)
	}()

	stmt, err := connection.Prepare(
		`SELECT 
    				id,
    				user_id,
    				friend_id,
    				status
			FROM friends WHERE user_id = $1 OR friend_id = $1`,
	)
	helpers.HandleError(err)

	rows, err := stmt.Query(userId)
	defer func(rows *sql.Rows) {
		helpers.HandleError(rows.Close())
	}(rows)

	var friends []friend.Friend

	for rows.Next() {
		var f friend.Friend
		err := rows.Scan(&f.Id, &f.UserId, &f.FriendId, &f.Status)
		helpers.HandleError(err)

		friends = append(friends, f)
	}

	c.JSON(200, gin.H{
		"data": friends,
	})
	return
}
