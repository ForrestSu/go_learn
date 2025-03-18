package p3

import "github.com/ForrestSu/go_learn/new-features/generic/p2"

// CheckIdentity checks the Identity method if it exists.
// Note that although this function calls a parameterized method,
// this function is not itself parameterized.
func CheckIdentity(v interface{}) {
	if vi, ok := v.(p2.HasIdentity); ok {
		if got := vi.Identity[int](0); got != 0 {
			panic(got)
		}
	}
}
