package specification

import "github.com/nullexp/specq/query"

type Executer interface {
	Execute(app Executable) any
}

type Specification interface {
	Executer
	And(Specification) Specification
	GroupAnd(Specification) Specification
	GroupOr(Specification) Specification
	Or(Specification) Specification
	Not() Specification
}

type CompositeSpecification struct {
	query query.Query
}

func NewSpecification(query query.Query) Specification {
	return CompositeSpecification{query: query}
}

func NewQuerySpecification(name string, op query.QueryOperator, operand *query.Operand) Specification {
	return CompositeSpecification{query: query.NewQuery(name, op, operand)}
}

func NewSpecificationWithModel(name string, op query.QueryOperator, model query.Identity, fields []string, operand *query.Operand) Specification {
	return CompositeSpecification{query: query.NewQueryWithModel(name, op, model, fields, operand)}
}

func NewEmptySpecification(name string, op query.QueryOperator) Specification {
	return CompositeSpecification{query: query.NewQuery(name, op, query.NewOperand(nil))}
}

func (c CompositeSpecification) Execute(applier Executable) any {
	return applier.Execute(c.query)
}

func (c CompositeSpecification) And(other Specification) Specification {
	return NewAndSpecification(c, other)
}

func (c CompositeSpecification) Or(left Specification) Specification {
	return NewOrSpecification(c, left)
}

func (c CompositeSpecification) GroupAnd(other Specification) Specification {
	return NewGroupAndSpecification(c, other)
}

func (c CompositeSpecification) Not() Specification {
	return NewNotSpecification(c)
}

func (c CompositeSpecification) GroupOr(other Specification) Specification {
	return NewGroupOrSpecification(c, other)
}
