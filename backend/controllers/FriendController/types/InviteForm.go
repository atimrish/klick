package types

import (
	"backend/database/db"
	"backend/helpers"
	"backend/validator"
)

type InviteForm struct {
	UserId   int64
	FriendId int64  `json:"friend_id"`
	Status   string
}

func (form *InviteForm) Validate() ([]string, bool) {
	hasError := false

	var messages []string

	validator.Required(form.UserId, "user_id", &messages)
	validator.Required(form.FriendId, "friend_id", &messages)
	validator.Required(form.Status, "status", &messages)

	if len(messages) != 0 {
		return messages, true
	}

	unique := !db.CheckUniquePostgres("users", "id", form.UserId)
	if unique {
		messages = append(messages, "user_id: такого пользователя не существует")
	}

	unique = !db.CheckUniquePostgres("users", "id", form.FriendId)
	if unique {
		messages = append(messages, "friend_id: такого пользователя не существует")
	}

	if len(messages) != 0 {
		return messages, true
	}

	if form.UserId == form.FriendId {
		messages = append(messages, "user_id и friend_id не могут быть равны")
		hasError = true
	}

	statuses := []string{"WAITING", "ACCEPTED", "DECLINED"}

	if !helpers.ArrayContains(&statuses, form.Status) {
		messages = append(messages, "неверное значение для поля 'status'")
		hasError = true
	}

	return messages, hasError
}
