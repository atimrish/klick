package main

import (
	"backend/cmd"
	"backend/database/db"
	"backend/helpers"
	"flag"
	"fmt"
)

func main() {
	action := flag.String("action", "", "action to do")
	flag.Parse()

	switch *action {

	case "serve":
		cmd.Serve()
		break
	case "migrate":
		cmd.Migrate()
		break

	case "migrate_down":
		cmd.Down()
		break

	case "mongo_ping":
		connection, ctx := db.MongoConnection()
		err := connection.Ping(ctx, nil)

		if err != nil {
			panic(err)
		}

		fmt.Println("successfully connect to mongo")
		break

	case "test":

		claims := helpers.GetPayloadJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDM4ODA0ODcsInRva2VuX2lkZW50aXR5Ijoi77-9IiwidXNlcl9pZCI6N30.EgPIsOEkhgXuJ78SN4uWnKT4qDnvv5wExWBQakxwaBI")
		fmt.Println(claims)

		break

	default:
		fmt.Printf("action doesn`t exists [%s]\n", *action)

	}

}
