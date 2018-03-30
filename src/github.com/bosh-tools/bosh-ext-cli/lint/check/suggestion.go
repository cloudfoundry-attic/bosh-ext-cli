package check

type Simple struct {
	Context_ Context

	Problem_    string
	Resolution_ string

	Details_ string
}

func (s Simple) Context() Context { return s.Context_ }
func (s Simple) Problem() string  { return s.Problem_ }

func (s Simple) Resolution() string {
	if len(s.Details_) > 0 {
		return s.Resolution_ + "\n" + "Details: " + s.Details_
	}
	return s.Resolution_
}
