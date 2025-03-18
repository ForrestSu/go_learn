package p4call

import (
	"github.com/ForrestSu/go_learn/new-features/generic/p1"
	"github.com/ForrestSu/go_learn/new-features/generic/p3"
)

// CheckSIdentity passes an S value to CheckIdentity.
func CheckSIdentity() {
	p3.CheckIdentity(p1.S{})
}
