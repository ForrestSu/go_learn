package main

import (
	"fmt"
)

type Receiver struct {
	Name string
}

// implements stringer
func (r *Receiver) String() string {
	return r.Name
}

type stringer func() string

func main() {
	methodValue() // 方法值
	useInMap()    // 在 map 中使用
}

func methodValue() {
	r := &Receiver{"Bob"}
	fn := r.String
	// (func() string)(0x1121920)
	fmt.Printf("%#v \n", fn)
}

// 在map中使用
func useInMap() {
	r := &Receiver{"Bob"}
	m := make(map[int]stringer)
	m[1] = r.String
	r.Name = "Alice"
	fmt.Printf("%s \n", m[1]())
}
