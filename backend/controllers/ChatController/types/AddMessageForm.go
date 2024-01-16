package types

type AddMessageForm struct {
	UserId int64  `json:"user_id"`
	Text   string `json:"text"`
}
