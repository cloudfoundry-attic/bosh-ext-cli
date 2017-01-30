package check

type Check interface {
	Description() Description
	Check() ([]Suggestion, error)
}

type Suggestion interface {
	Context() Context
	Problem() string
	Resolution() string
}
