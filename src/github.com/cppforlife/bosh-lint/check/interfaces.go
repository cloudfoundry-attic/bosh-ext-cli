package check

type Check interface {
	Description() Description
	Check() ([]Suggestion, error)
	IsEnabled() bool
}

type Suggestion interface {
	Context() Context
	Problem() string
	Resolution() string
}
