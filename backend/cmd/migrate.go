package cmd

import (
	"backend/database/helpers"
	"backend/database/migrations/postgres"
	"fmt"
)

var (
	migrationsUp = []func() string {
		postgres.UpUsersTable,
		postgres.UpFriendsTable,
	}

	migrationsDown = []func() string {
		postgres.DownUsersTable,
		postgres.DownFriendsTable,
	}
)

func Migrate() {
	connection := helpers.PostgresConnection()
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
	connection := helpers.PostgresConnection()
	defer connection.Close()

	for _, f := range migrationsDown {
		sql := f()

		_, err := connection.Exec(sql)

		if err != nil {
			fmt.Println("migrate down error")
			panic(err)
		}
	}

	fmt.Println("migrated down successfully")
}
