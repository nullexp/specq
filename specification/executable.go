package specification

import "github.com/nullexp/specq/query"

type Executable interface {
	Execute(query.Query) any
	And(left, right Specification) any
	Or(left, right Specification) any
	Not(spec Specification) any
	GroupAnd(left, right Specification) any
	GroupOr(left, right Specification) any
}
