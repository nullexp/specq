package specification

// AndSpecification
type GroupAndSpecficiation struct {
	rightCondition, leftCondition Specification
}

func NewGroupAndSpecification(leftCondition, rightCondition Specification) Specification {
	spec := GroupAndSpecficiation{leftCondition: leftCondition, rightCondition: rightCondition}
	return spec
}

func (c GroupAndSpecficiation) And(other Specification) Specification {
	return NewAndSpecification(c, other)
}

func (c GroupAndSpecficiation) GroupAnd(other Specification) Specification {
	return NewGroupAndSpecification(c, other)
}

func (c GroupAndSpecficiation) Not() Specification {
	return NewNotSpecification(c)
}

func (c GroupAndSpecficiation) Or(left Specification) Specification {
	return NewOrSpecification(c, left)
}

func (a GroupAndSpecficiation) Execute(app Executable) any {
	return app.GroupAnd(a.leftCondition, a.rightCondition)
}

func (c GroupAndSpecficiation) GroupOr(other Specification) Specification {
	return NewGroupOrSpecification(c, other)
}
