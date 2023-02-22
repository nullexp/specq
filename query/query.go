package query

type Query interface {
	GetName() string
	SetName(string)
	GetOperator() QueryOperator
	GetOperand() *Operand // The value/values which are used in the query
	GetModel() Identity   // The entity which the query is for
	GetFields() []string  // Wanted fields
}

type basicQuery struct {
	name    string
	op      QueryOperator
	operand *Operand
	model   Identity
	fields  []string
}

func NewQueryWithModel(name string, op QueryOperator, model Identity, fields []string, operand *Operand) Query {
	return &basicQuery{name: name, op: op, operand: operand, fields: fields, model: model}
}

func NewQuery(name string, op QueryOperator, operand *Operand) Query {
	return &basicQuery{name: name, op: op, operand: operand, fields: []string{}, model: nil}
}

func NewEmptyQuery(name string, op QueryOperator) Query {
	return &basicQuery{name: name, op: op, operand: nil, fields: []string{}, model: nil}
}

func (b basicQuery) GetModel() Identity {
	return b.model
}

func (b basicQuery) GetFields() []string {
	return b.fields
}

func (b basicQuery) GetName() string {
	return b.name
}

func (b *basicQuery) SetName(name string) {
	b.name = name
}

func (b basicQuery) GetOperator() QueryOperator {
	return b.op
}

func (b basicQuery) GetOperand() *Operand {
	return b.operand
}

func MapQueryWithFields(q Query, identity Identity, fields []string) Query {
	return NewQueryWithModel(q.GetName(), q.GetOperator(), identity, fields, q.GetOperand())
}

func MapQueryWithModel(q Query, identity Identity) Query {
	return MapQueryWithFields(q, identity, []string{})
}

func MapQueriesWithModel(qs []Query, identity Identity) []Query {
	return MapQueriesWithFields(qs, identity, []string{})
}

func MapQueries(qs []Query) []Query {
	return MapQueriesWithFields(qs, nil, []string{})
}

func MapQueriesWithFields(qs []Query, identity Identity, fields []string) (out []Query) {
	out = []Query{}
	for _, v := range qs {
		out = append(out, MapQueryWithFields(v, identity, fields))
	}
	return out
}

type Operand struct {
	Value any
}

func NewOperand(value any) *Operand { return &Operand{Value: value} }
