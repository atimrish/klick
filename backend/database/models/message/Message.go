package message

import (
	"backend/database/models/post"
	"time"
)

const tableName = "messages"

type Message struct {
	Id          int64        `bson:"id,omitempty"`
	UserId      int64        `bson:"user_id,omitempty"`
	Text        string       `bson:"text,omitempty"`
	Photos      []string     `bson:"photos,omitempty"`
	Videos      []post.Video `bson:"videos,omitempty"`
	CreatedTime time.Time    `bson:"created_time"`
	UpdatedTime time.Time    `bson:"updated_time"`
}
