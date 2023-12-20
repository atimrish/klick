package main

import (
	"backend/cmd"
	"backend/database/db"
	"backend/database/models/user"
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

	case "create_user":
		tmpUser := user.NewUser(0, "test", "test", "test", "test", "test", "test")

		err := user.CreateUser(tmpUser)
		helpers.HandleError(err)

		fmt.Println("user created")
		break

	case "get_users":
		fmt.Println(user.GetAllUsers())
		break

	case "update_user":
		users := user.GetAllUsers()
		tmpUser := (*users)[0]
		user.SetUserName(&tmpUser, "updated")
		user.UpdateUser(&tmpUser)

		fmt.Println("user updated")
		break

	case "delete_user":
		users := user.GetAllUsers()
		tmpUser := (*users)[0]
		user.DeleteUser(user.GetUserId(&tmpUser))

		fmt.Println("user deleted")
		break

	default:
		fmt.Printf("action doesn`t exists [%s]\n", *action)

	}

}
