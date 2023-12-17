package main

import (
	"backend/cmd"
	"backend/database/helpers"
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
		connection, ctx := helpers.MongoConnection()
		err := connection.Ping(ctx, nil)

		if err != nil {
			panic(err)
		}

		fmt.Println("successfully connect to mongo")
		break

	default:
		fmt.Printf("action doesn`t exists [%s]\n", *action)

	}

}
