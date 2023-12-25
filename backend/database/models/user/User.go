package user

import (
	"backend/database/db"
	"backend/helpers"
	"fmt"
)

const tableName = "users"

type User struct {
	Id       int64
	Surname  string
	Name     string
	Login    string
	Password string
	Email    string
	Photo    string
}

func NewUser(id int64, surname, name, login, password, email, photo string) *User {
	tmpUser := User{
		Id:       id,
		Surname:  surname,
		Name:     name,
		Login:    login,
		Password: password,
		Email:    email,
		Photo:    photo,
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
		err := rows.Scan(&user.Id, &user.Surname, &user.Name, &user.Login, &user.Password, &user.Email, &user.Photo)
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

	query = fmt.Sprintf(query, user.Surname, user.Name, user.Login, user.Password, user.Email, user.Photo)
	connection.QueryRow(query).Scan(&user.Id)

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

	query = fmt.Sprintf(query, user.Surname, user.Name, user.Login, user.Email, user.Photo, user.Id)

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

func FindUserByLogin(login string) (User, bool) {
	connection := db.PostgresConnection()
	defer connection.Close()

	query := `SELECT id, surname, name, login, password, email, photo FROM users WHERE login = '%s'`
	rows, err := connection.Query(fmt.Sprintf(query, login))

	helpers.HandleError(err)
	defer rows.Close()

	var user User
	hasUser := false
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Surname, &user.Name, &user.Login, &user.Password, &user.Email, &user.Photo)
		helpers.HandleError(err)
		hasUser = true
	}

	return user, hasUser
}
