package message

import (
	"backend/database/models/post"
	"time"
)

const tableName = "messages"

type Message struct {
	id          uint16
	userId      uint16
	text        string
	photos      []string
	videos      []post.Video
	createdTime time.Time
	updatedTime time.Time
}
