package task

import (
	"time"
)

type Line struct {
	Number  int
	Level   string
	Time    time.Time
	Group   string
	Content string

	action Action
}

type UnknownLine struct {
	Number  int
	Content string
}

func (l Line) IsError() bool { return l.Level == "E" }

func (l Line) TruncatedContent(max int) string {
	len := len(l.Content)
	if len == 0 {
		return "<empty>"
	}
	if len > max {
		return l.Content[0:max] + " ..."
	}
	return l.Content
}

func (l *Line) Action() Action {
	if l.action != nil {
		return l.action
	}

	if a := NewAgentRequest(l.Content); a != nil {
		l.action = a
	} else if a := NewAgentResponse(l.Content); a != nil {
		l.action = a
	} else if a := NewCPIRequest(l.Content); a != nil {
		l.action = a
	} else if a := NewCPIResponse(l.Content); a != nil {
		l.action = a
	} else if a := NewDBStatement(l.Content, l.Number); a != nil {
		l.action = a
	} else {
		l.action = UnknownAction{l.Number, l.TruncatedContent(80)}
	}

	return l.action
}

type LineTimeSorting []*Line

func (ls LineTimeSorting) Len() int { return len(ls) }

func (ls LineTimeSorting) Less(i, j int) bool {
	if ls[i].Time.Equal(ls[j].Time) {
		return ls[i].Number < ls[j].Number
	}
	return ls[i].Time.Before(ls[j].Time)
}

func (ls LineTimeSorting) Swap(i, j int) { ls[i], ls[j] = ls[j], ls[i] }
