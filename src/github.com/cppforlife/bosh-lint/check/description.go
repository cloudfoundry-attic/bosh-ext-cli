package check

type Description struct {
	Context_   Context
	Purpose_   string
	Reasoning_ string
}

func (d Description) Context() Context { return d.Context_ }
func (d Description) Purpose() string  { return "checks " + d.Purpose_ }
