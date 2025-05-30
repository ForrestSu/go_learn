package p1

// S  is  a  type  with  a  parameterized  method  Identity.
type S struct{}

// Identity  is  a  simple  identity  method  that  works  for  any  type.
func (S) Identity[T any](v T) T { return v }
