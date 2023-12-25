package types

import (
	"backend/database/db"
	"mime/multipart"
)

type RegisterForm struct {
	Surname  string `form:"surname" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Login    string `form:"login" binding:"required"`
	Password string `form:"password" binding:"required"`
	Photo    *multipart.FileHeader `form:"photo"`
}

func (form RegisterForm) Validate() ([]string, bool) {
	hasError := false

	var messages []string

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
