package types

import (
	"backend/database/db"
	"backend/validator"
	"mime/multipart"
)

type RegisterForm struct {
	Surname  string                `form:"surname"`
	Name     string                `form:"name"`
	Email    string                `form:"email"`
	Login    string                `form:"login"`
	Password string                `form:"password"`
	Photo    *multipart.FileHeader `form:"photo"`
}

func (form RegisterForm) Validate() ([]string, bool) {
	hasError := false

	var messages []string

	validator.Required(form.Surname, "surname", &messages)
	validator.Required(form.Name, "name", &messages)
	validator.Required(form.Email, "email", &messages)
	validator.Required(form.Login, "login", &messages)
	validator.Required(form.Password, "password", &messages)

	if len(messages) != 0 {
		return messages, true
	}

	if len(form.Password) < 8 {
		hasError = true
		messages = append(messages, "Количество символов в пароле должно быть не меньше 8")
	}

	if !db.CheckUniquePostgres("users", "email", form.Email) {
		hasError = true
		messages = append(messages, "Такой email уже занят")
	}

	if !db.CheckUniquePostgres("users", "login", form.Login) {
		hasError = true
		messages = append(messages, "Такой логин уже занят")
	}

	return messages, hasError
}
