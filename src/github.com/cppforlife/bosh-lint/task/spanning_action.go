package task

import (
	"fmt"
	"time"
)

type SpanningAction struct {
	Lines []*Line
}

func (a SpanningAction) StartedAt() time.Time    { return a.Lines[0].Time }
func (a SpanningAction) EndedAt() time.Time      { return a.Lines[len(a.Lines)-1].Time }
func (a SpanningAction) Duration() time.Duration { return a.EndedAt().Sub(a.StartedAt()) }

func (a SpanningAction) ShortDescription() string {
	first := a.Lines[0].Action().ShortDescription()
	len := len(a.Lines)
	if len == 1 {
		return first + " 1"
	}
	last := a.Lines[len-1].Action().ShortDescription()
	return fmt.Sprintf("%s -> %s %d", first, last, len)
}

func (a *SpanningAction) isRelated(b *SpanningAction) bool {
	for _, lineA := range a.Lines {
		relA := lineA.Action().Relation()

		for _, lineB := range b.Lines {
			relB := lineB.Action().Relation()

			if relA.Matches(relB) {
				return true
			}
		}
	}

	return false
}
