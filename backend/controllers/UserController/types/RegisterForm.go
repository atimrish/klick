package types

type RegisterForm struct {
	Surname  string `json:"surname"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}
