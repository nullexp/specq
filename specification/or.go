package specification

type OrSpecficiation struct {
	leftCondition, rightCondition Specification
}

func NewOrSpecification(leftCondition, rightCondition Specification) Specification {
	spec := OrSpecficiation{leftCondition: leftCondition, rightCondition: rightCondition}
	return spec
}

func (c OrSpecficiation) And(other Specification) Specification {
	return NewAndSpecification(c, other)
}

func (c OrSpecficiation) Or(left Specification) Specification {
	return NewOrSpecification(c, left)
}

func (c OrSpecficiation) GroupAnd(other Specification) Specification {
	return NewGroupAndSpecification(c, other)
}

func (c OrSpecficiation) Not() Specification {
	return NewNotSpecification(c)
}

func (a OrSpecficiation) Execute(app Executable) any {
	return app.Or(a.leftCondition, a.rightCondition)
}

func (c OrSpecficiation) GroupOr(other Specification) Specification {
	return NewGroupOrSpecification(c, other)
}
