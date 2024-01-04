package friend

import (
	"backend/database/db"
	"backend/helpers"
)

const tableName = "friends"

type Friend struct {
	Id       int64
	UserId   int64
	FriendId int64
	Status   string
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
