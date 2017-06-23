package task

type Action interface {
	Relation() Relation
	ShortDescription() string
}

type UnknownAction struct {
	description string
}

func (a UnknownAction) Relation() Relation       { return ExactRelation{""} }
func (a UnknownAction) ShortDescription() string { return a.description }

type Relation interface {
	Matches(Relation) bool
}

type ExactRelation struct {
	Value string
}

var _ Relation = ExactRelation{}

func (id1 ExactRelation) Matches(id Relation) bool {
	if id2, ok := id.(ExactRelation); ok {
		return id1.Value == id2.Value
	}
	return false
}

type NonMatchingRelation struct{}

var _ Relation = NonMatchingRelation{}

func (id1 NonMatchingRelation) Matches(id Relation) bool { return false }
