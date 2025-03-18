package p2

// HasIdentity is an interface that matches any type with a
// parameterized Identity method.
type HasIdentity interface {
	Identity[T any](T) T
}
