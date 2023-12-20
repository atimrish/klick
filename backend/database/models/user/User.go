package user

import (
	"backend/database/db"
	"backend/helpers"
	"fmt"
)

const tableName = "users"

type User struct {
	id       int64
	surname  string
	name     string
	login    string
	password string
	email    string
	photo    string
}

func NewUser(id int64, surname, name, login, password, email, photo string) *User {
	tmpUser := User{
		id:       id,
		surname:  surname,
		name:     name,
		login:    login,
		password: password,
		email:    email,
		photo:    photo,
	}

	return &tmpUser
}

func GetAllUsers() *[]User {
	connection := db.PostgresConnection()
	defer connection.Close()

	rows, err := connection.Query(`SELECT id, surname, name, login, password, email, photo FROM users`)
	helpers.HandleError(err)
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.id, &user.surname, &user.name, &user.login, &user.password, &user.email, &user.photo)
		helpers.HandleError(err)

		users = append(users, user)
	}

	return &users
}

func CreateUser(user *User) error {
	connection := db.PostgresConnection()
	defer connection.Close()

	query := `INSERT INTO users(surname, name, login, password, email, photo) 
				VALUES ('%s', '%s', '%s', '%s', '%s', '%s') RETURNING id`

	query = fmt.Sprintf(query, user.surname, user.name, user.login, user.password, user.email, user.photo)
	connection.QueryRow(query).Scan(&user.id)

	return nil
}

func UpdateUser(user *User) {
	connection := db.PostgresConnection()
	defer connection.Close()

	query := `
		UPDATE users
			SET
			    surname = '%s',
			    name = '%s',
			    login = '%s',
			    email = '%s',
			    photo = '%s',
				updated_at = CURRENT_TIMESTAMP
		WHERE users.id = %d`

	query = fmt.Sprintf(query, user.surname, user.name, user.login, user.email, user.photo, user.id)

	_, err := connection.Exec(query)
	helpers.HandleError(err)
}

func DeleteUser(id int64) {
	connection := db.PostgresConnection()
	defer connection.Close()

	query := `DELETE FROM users WHERE users.id = %d`
	query = fmt.Sprintf(query, id)

	_, err := connection.Exec(query)
	helpers.HandleError(err)
}

func SetUserName(user *User, name string) {
	user.name = name
}

func GetUserId(user *User) int64 {
	return user.id
}
