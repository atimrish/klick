package message

import (
	"backend/database/db"
	"backend/database/models/post"
	"backend/helpers"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

const database = "klick"
const tableName = "chats"

type Message struct {
	Id          string       `bson:"id,omitempty" json:"id,omitempty"`
	UserId      int64        `bson:"user_id,omitempty" json:"user_id"`
	Text        string       `bson:"text,omitempty" json:"text"`
	Photos      []string     `bson:"photos,omitempty" json:"photos,omitempty"`
	Videos      []post.Video `bson:"videos,omitempty" json:"videos,omitempty"`
	CreatedTime string       `bson:"created_time" json:"created_time"`
	UpdatedTime string       `bson:"updated_time" json:"updated_time"`
}

func (m *Message) PushMessage(chatId primitive.ObjectID) {
	client, context := db.MongoConnection()

	filter := bson.D{
		{
			"_id",
			chatId,
		},
	}

	m.Id = uuid.New().String()
	m.CreatedTime = strconv.Itoa(int(time.Now().Unix()))

	update := bson.M{
		"$push": bson.M{
			"messages": m,
		},
	}

	_, err := client.Database(database).Collection(tableName).UpdateOne(context, filter, update)
	helpers.HandleError(err)

	return
}

func UpdateMessage(chatId primitive.ObjectID, messageId, text string) {
	client, context := db.MongoConnection()

	filter := bson.M{
		"_id":         chatId,
		"messages.id": messageId,
	}
	update := bson.M{
		"$set": bson.M{
			"messages.$[].text": text,
		},
	}

	_, err := client.Database(database).Collection(tableName).UpdateOne(context, filter, update)
	helpers.HandleError(err)

	return
}
