package db

import (
	"backend/conf"
	"backend/helpers"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PostgresConnection() *sql.DB {
	config := conf.GetConfig()

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config["postgres_address"],
		config["postgres_port"],
		config["postgres_user"],
		config["postgres_password"],
		config["postgres_database"],
	)

	connection, err := sql.Open("postgres", connectionString)
	helpers.HandleError(err)

	err = connection.Ping()
	helpers.HandleError(err)

	return connection
}

func CheckUniquePostgres(table, column string, value any) bool {
	connection := PostgresConnection()
	defer connection.Close()
	query := `SELECT %s as needle FROM %s`
	query = fmt.Sprintf(query, column, table)

	rows, err := connection.Query(query)
	defer rows.Close()
	helpers.HandleError(err)

	for rows.Next() {
		var tmpVal string
		err := rows.Scan(&tmpVal)
		fmt.Println("log: ", tmpVal)
		fmt.Println("log [=] : ", tmpVal == value)
		helpers.HandleError(err)

		if value == tmpVal {
			return false
		}
	}

	return true
}

func MongoConnection() (*mongo.Client, context.Context) {
	config := conf.GetConfig()

	address := fmt.Sprintf("mongodb://%s:%s/", config["mongo_address"], config["mongo_post"])

	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(address)
	client, err := mongo.Connect(ctx, clientOptions)
	helpers.HandleError(err)

	return client, ctx
}
