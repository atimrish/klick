package chat

import (
	models2 "backend/database/models/message"
)

const tableName = "chats"

type Chat struct {
	id       uint16
	users    []int16
	messages []models2.Message
	photo    string
}
