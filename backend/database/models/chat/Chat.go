package chat

import (
	"backend/database/db"
	models "backend/database/models/message"
	"backend/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const tableName = "chats"
const database = "klick"

type Chat struct {
	Id       primitive.ObjectID `bson:"id,omitempty" json:"id"`
	Users    []int64            `bson:"users" json:"users"`
	Title    string             `bson:"title" json:"title"`
	Messages []models.Message   `bson:"messages" json:"messages"`
	Photo    string             `bson:"photo" json:"photo"`
}

func (c *Chat) Insert() {
	client, context := db.MongoConnection()

	res, err := client.Database(database).Collection(tableName).InsertOne(context, c)
	helpers.HandleError(err)
	c.Id = res.InsertedID.(primitive.ObjectID)

	return
}
