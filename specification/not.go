package specification

type NotSpecficiation struct {
	wrapped Specification
}

func NewNotSpecification(condition Specification) Specification {
	spec := NotSpecficiation{wrapped: condition}
	return spec
}

func (c NotSpecficiation) And(other Specification) Specification {
	return NewAndSpecification(c, other)
}

func (c NotSpecficiation) Not() Specification {
	return NewNotSpecification(c.wrapped)
}

func (c NotSpecficiation) Or(left Specification) Specification {
	return NewOrSpecification(c, left)
}

func (c NotSpecficiation) GroupAnd(other Specification) Specification {
	return NewGroupAndSpecification(c, other)
}

func (a NotSpecficiation) Execute(app Executable) any {
	return app.Not(a.wrapped)
}

func (c NotSpecficiation) GroupOr(other Specification) Specification {
	return NewGroupOrSpecification(c, other)
}
