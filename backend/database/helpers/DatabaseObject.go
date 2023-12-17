package helpers

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PostgresConnection() *sql.DB  {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		5432,
		"root",
		"root",
		"social_network",
		)

	connection, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	err = connection.Ping()
	if err != nil {
		panic(err)
	}

	return connection
}

func MongoConnection() (*mongo.Client, context.Context) {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	return client, ctx
}
