package cmd

import (
	"backend/database/db"
	"backend/database/migrations/postgres"
	"backend/helpers"
	"fmt"
)

var (
	migrationsUp = []func() string{
		postgres.UpUsersTable,
		postgres.UpFriendsTable,
	}

	migrationsDown = []func() string{
		postgres.DownUsersTable,
		postgres.DownFriendsTable,
	}
)

func Migrate() {
	connection := db.PostgresConnection()
	defer connection.Close()

	for _, f := range migrationsUp {
		sql := f()

		_, err := connection.Exec(sql)

		if err != nil {
			fmt.Println("migrate error")
			panic(err)
		}
	}

	fmt.Println("migrated successfully")
}

func Down() {
	connection := db.PostgresConnection()
	defer connection.Close()

	migrations := helpers.ArrayReverse(&migrationsDown)
	
	for _, f := range migrations {
		sql := f()

		_, err := connection.Exec(sql)

		if err != nil {
			fmt.Println("migrate down error")
			panic(err)
		}
	}

	fmt.Println("migrated down successfully")
}
