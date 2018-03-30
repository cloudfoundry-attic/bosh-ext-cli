package check

type Context struct {
	parent      *Context
	description string
}

func NewRootContext(desc string) Context {
	return Context{nil, desc}
}

func (c Context) Nested(desc string) Context {
	return Context{&c, desc}
}

func (c Context) String() string {
	if c.parent == nil {
		return c.description
	}
	return c.parent.String() + ": " + c.description
}
