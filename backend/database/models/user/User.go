package user

import "backend/database/helpers"

const tableName = "users"

type User struct {
	id       uint16
	surname  string
	name     string
	login    string
	password string
	email    string
	photo    string
}

func GetAllUsers() []User {
	db := helpers.PostgresConnection()
	db.Ping()

	return []User{}
}

