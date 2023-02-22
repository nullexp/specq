package specification

// AndSpecification
type AndSpecficiation struct {
	rightCondition, leftCondition Specification
}

func NewAndSpecification(leftCondition, rightCondition Specification) Specification {
	spec := AndSpecficiation{leftCondition: leftCondition, rightCondition: rightCondition}

	return spec
}

func (c AndSpecficiation) GroupAnd(other Specification) Specification {
	return NewGroupAndSpecification(c, other)
}

func (c AndSpecficiation) And(other Specification) Specification {
	return NewAndSpecification(c, other)
}

func (c AndSpecficiation) Not() Specification {
	return NewNotSpecification(c)
}

func (c AndSpecficiation) Or(left Specification) Specification {
	return NewOrSpecification(c, left)
}

func (a AndSpecficiation) Execute(app Executable) any {
	return app.And(a.leftCondition, a.rightCondition)
}

func (c AndSpecficiation) GroupOr(other Specification) Specification {
	return NewGroupOrSpecification(c, other)
}
