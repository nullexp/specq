package query

// Identity represent an entity which can be any type but we use int64 for simplicity
type Identity interface {
	GetID() int64
}
