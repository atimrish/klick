package friend

const tableName = "friends"

type Friend struct {
	Id       int64
	UserId   int64
	FriendId int64
	Status   string
}
