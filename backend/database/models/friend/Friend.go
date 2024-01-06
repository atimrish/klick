package friend

import (
	"backend/database/db"
	"backend/helpers"
)

const tableName = "friends"

type Friend struct {
	Id       int64  `json:"id,omitempty"`
	UserId   int64  `json:"user_id,omitempty"`
	FriendId int64  `json:"friend_id,omitempty"`
	Status   string `json:"status,omitempty"`
}

func (f *Friend) Update() {
	connection := db.PostgresConnection()

	defer func() {
		err := connection.Close()
		helpers.HandleError(err)
	}()

	stmt, err := connection.Prepare(`UPDATE friends SET user_id = $1, friend_id = $2, status = $3 WHERE id = $4`)
	helpers.HandleError(err)

	_, err = stmt.Exec(f.UserId, f.FriendId, f.Status, f.Id)
	helpers.HandleError(err)

	return
}
