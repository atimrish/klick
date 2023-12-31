package validator

import "fmt"

func Required(value any, title string, messages *[]string) {
	if value == "" {
		*messages = append(*messages, fmt.Sprintf("Поле '%s' должно быть заполнено", title))
	}
}
