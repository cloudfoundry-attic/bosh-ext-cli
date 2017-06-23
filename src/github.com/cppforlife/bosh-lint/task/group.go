package task

import (
	"sort"
	"time"
)

type Group struct {
	Name  string
	Lines []*Line
}

func (g Group) Actions() []SpanningAction {
	var actions []*SpanningAction

	for _, line := range g.Lines {
		actions = append(actions, &SpanningAction{Lines: []*Line{line}})
	}

	for {
		var mergedOnce bool

		for i1, action1 := range actions {
			if action1 == nil {
				continue
			}
			for i2, action2 := range actions {
				if action2 == nil {
					continue
				}
				if i1 != i2 && action1.isRelated(action2) {
					actions[i1].Lines = append(actions[i1].Lines, action2.Lines...)
					actions[i2] = nil
					mergedOnce = true
				}
			}
		}

		if !mergedOnce {
			break
		}
	}

	var actions2 []SpanningAction

	for _, action := range actions {
		if action != nil {
			sort.Sort(LineTimeSorting(action.Lines))
			actions2 = append(actions2, *action)
		}
	}

	return actions2
}

func (g Group) StartedAt() time.Time    { return g.Lines[0].Time }
func (g Group) EndedAt() time.Time      { return g.Lines[len(g.Lines)-1].Time }
func (g Group) Duration() time.Duration { return g.EndedAt().Sub(g.StartedAt()) }
