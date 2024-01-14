package types

import (
	"backend/helpers"
	"strconv"
	"strings"
)

type CreateChatForm struct {
	Users string `json:"users"`
}

func (f *CreateChatForm) GetUsersArray() *[]int64 {
	users := strings.Split(f.Users, ",")
	var returnedArray []int64

	for _, user := range users {
		userInt, err := strconv.Atoi(user)
		helpers.HandleError(err)
		returnedArray = append(returnedArray, int64(userInt))
	}

	return &returnedArray
}
