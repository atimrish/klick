package chat

import (
	"backend/database/db"
	models "backend/database/models/message"
	"backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	tableName    = "chats"
	database     = "klick"
	TYPE_PRIVATE = "PRIVATE"
	TYPE_GROUP   = "GROUP"
)

type Chat struct {
	Id       primitive.ObjectID `bson:"id,omitempty" json:"id"`
	Users    []int64            `bson:"users" json:"users"`
	Title    string             `bson:"title" json:"title"`
	Messages []models.Message   `bson:"messages" json:"messages"`
	Photo    string             `bson:"photo" json:"photo"`
	Type     string             `bson:"type" json:"type"`
}

func GetByUserId(userId int64) (*[]Chat, error) {
	client, context := db.MongoConnection()

	filter := bson.M{
		"users": userId,
	}

	cursor, err := client.Database(database).Collection(tableName).Find(context, filter)
	helpers.HandleError(err)

	var chats []Chat
	err = cursor.All(context, &chats)

	return &chats, err
}

func GetById(chatId string) (*Chat, error) {
	objId, err := primitive.ObjectIDFromHex(chatId)
	helpers.HandleError(err)

	client, context := db.MongoConnection()
	filter := bson.D{{
		"_id", objId,
	}}

	res := client.Database(database).Collection(tableName).FindOne(context, filter)
	var returnedChat Chat
	err = res.Decode(&returnedChat)

	return &returnedChat, err
}

func (c *Chat) Insert() {
	client, context := db.MongoConnection()

	res, err := client.Database(database).Collection(tableName).InsertOne(context, c)
	helpers.HandleError(err)
	c.Id = res.InsertedID.(primitive.ObjectID)

	return
}
