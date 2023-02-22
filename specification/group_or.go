package specification

// AndSpecification
type GrouporSpecficiation struct {
	rightCondition, leftCondition Specification
}

func NewGroupOrSpecification(leftCondition, rightCondition Specification) Specification {
	spec := GrouporSpecficiation{leftCondition: leftCondition, rightCondition: rightCondition}
	return spec
}

func (c GrouporSpecficiation) And(other Specification) Specification {
	return NewAndSpecification(c, other)
}

func (c GrouporSpecficiation) GroupAnd(other Specification) Specification {
	return NewGroupAndSpecification(c, other)
}

func (c GrouporSpecficiation) GroupOr(other Specification) Specification {
	return NewGroupOrSpecification(c, other)
}

func (c GrouporSpecficiation) Not() Specification {
	return NewNotSpecification(c)
}

func (c GrouporSpecficiation) Or(left Specification) Specification {
	return NewOrSpecification(c, left)
}

func (a GrouporSpecficiation) Execute(app Executable) any {
	return app.GroupOr(a.leftCondition, a.rightCondition)
}
