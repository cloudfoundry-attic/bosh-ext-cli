package task

import (
	"fmt"
	"time"

	boshtbl "github.com/cloudfoundry/bosh-cli/ui/table"
)

type ValueDuration struct {
	D time.Duration
}

func NewValueDuration(d time.Duration) ValueDuration { return ValueDuration{D: d} }

func (t ValueDuration) String() string       { return fmt.Sprintf("%s", t.D) }
func (t ValueDuration) Value() boshtbl.Value { return t }

func (t ValueDuration) Compare(other boshtbl.Value) int {
	otherD := other.(ValueDuration)
	switch {
	case t.D == otherD.D:
		return 0
	case t.D < otherD.D:
		return -1
	default:
		return 1
	}
}
