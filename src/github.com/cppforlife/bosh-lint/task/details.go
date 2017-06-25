package task

import (
	"fmt"
)

type Details struct {
	DirectorVersion string
	TaskID          string

	Lines        []*Line
	UnknownLines []UnknownLine
}

func (d Details) FindGroup(name string) (Group, error) {
	if len(name) == 0 {
		return d.Group(), nil
	}

	for _, group := range d.Groups() {
		if group.Name == name {
			return group, nil
		}
	}

	return Group{}, fmt.Errorf("Did not find group '%s'", name)
}

func (d Details) Group() Group {
	return Group{"", d.Lines}
}

func (d Details) Groups() []Group {
	groupsByName := map[string][]*Line{}

	for _, line := range d.Lines {
		// todo group limits nats responses
		groupsByName[line.Group] = append(groupsByName[line.Group], line)
	}

	var groups []Group

	for name, lines := range groupsByName {
		groups = append(groups, Group{name, lines})
	}

	return groups
}
