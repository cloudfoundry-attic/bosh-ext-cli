package check

import (
	"strings"
)

type Todo struct {
	context Context
	content string
	Config
}

func NewTodo(context Context, content string, config Config) Todo {
	return Todo{context, content, config}
}

func (c Todo) Description() Description {
	return Description{
		Context_: c.context,
		Purpose_: "if it's a todo",
	}
}

func (c Todo) Check() ([]Suggestion, error) {
	var sugs []Suggestion

	if strings.Contains(strings.ToLower(c.content), "todo") {
		sugs = append(sugs, Simple{
			Context_:    c.context,
			Problem_:    "Todo",
			Resolution_: "Do",
		})
	}

	return sugs, nil
}
