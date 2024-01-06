package post

import (
	"backend/database/db"
	"backend/helpers"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const tableName = "posts"
const database = "klick"

type Post struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Text       string             `bson:"text" json:"text"`
	Photos     []string           `bson:"photos" json:"photos"`
	UserId     int64              `bson:"userId" json:"user_id"`
	Videos     []Video            `bson:"videos" json:"videos"`
	Categories []int64            `bson:"categories" json:"categories"`
}

type Video struct {
	Link    string `bson:"link" json:"link"`
	Preview string `bson:"preview" json:"preview"`
}

func (p *Post) Insert() {
	client, context := db.MongoConnection()
	res, err := client.Database(database).Collection(tableName).InsertOne(context, p)
	helpers.HandleError(err)
	p.Id = res.InsertedID.(primitive.ObjectID)
}

func GetPosts(limit, offset int64) (*[]Post, error) {
	client, context := db.MongoConnection()

	cursor, err := client.Database(database).
		Collection(tableName).
		Find(context, bson.D{}, options.Find().
			SetLimit(limit).
			SetSkip(offset),
		)

	helpers.HandleError(err)
	var posts []Post
	err = cursor.All(context, &posts)

	return &posts, err
}

func GetPostById(id string) (Post, error) {
	client, context := db.MongoConnection()

	objId, err := primitive.ObjectIDFromHex(id)
	helpers.HandleError(err)

	filter := bson.D{{
		"_id", objId,
	}}

	var post Post

	res := client.Database(database).Collection(tableName).FindOne(context, filter)
	fmt.Println("res ", res)

	err = res.Decode(&post)
	fmt.Println("error", err)

	return post, err
}

func DeletePostById(id string) error {
	client, context := db.MongoConnection()

	objId, err := primitive.ObjectIDFromHex(id)
	helpers.HandleError(err)

	filter := bson.D{{
		"_id", objId,
	}}

	_, err = client.Database(database).Collection(tableName).DeleteOne(context, filter)

	return err
}

func (p *Post) Update() error {
	client, context := db.MongoConnection()
	_, err := client.Database(database).Collection(tableName).UpdateOne(context, p.Id, p)
	return err
}
