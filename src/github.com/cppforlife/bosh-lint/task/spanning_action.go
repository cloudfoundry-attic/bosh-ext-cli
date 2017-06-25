package task

import (
	"fmt"
	"sort"
	"time"
)

type SpanningAction struct {
	lines           []*Line
	uniqueRelations []Relation
}

func (a SpanningAction) StartedAt() time.Time    { return a.lines[0].Time }
func (a SpanningAction) EndedAt() time.Time      { return a.lines[len(a.lines)-1].Time }
func (a SpanningAction) Duration() time.Duration { return a.EndedAt().Sub(a.StartedAt()) }

func (a SpanningAction) ShortDescription() string {
	first := a.lines[0].Action().ShortDescription()
	len := len(a.lines)
	if len == 1 {
		return first + " 1"
	}
	last := a.lines[len-1].Action().ShortDescription()
	return fmt.Sprintf("%s -> %s %d", first, last, len)
}

func (a *SpanningAction) IsRelated(b *SpanningAction) bool {
	for _, relA := range a.uniqueRelations {
		for _, relB := range b.uniqueRelations {
			if relA.Matches(relB) {
				return true
			}
		}
	}

	return false
}

func (a *SpanningAction) Merge(b *SpanningAction) {
	a.AddLines(b.lines)
}

func (a *SpanningAction) AddLines(lines []*Line) {
	for _, line := range lines {
		lineRel := line.Action().Relation()
		var found bool

		for _, rel := range a.uniqueRelations {
			if lineRel == rel {
				found = true
				break
			}
		}

		if !found {
			a.uniqueRelations = append(a.uniqueRelations, lineRel)
		}
	}

	a.lines = append(a.lines, lines...)
}

func (a *SpanningAction) SortLinesChrono() {
	sort.Sort(LineTimeSorting(a.lines))
}
