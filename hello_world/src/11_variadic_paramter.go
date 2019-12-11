package main

import "fmt"

func change(s ...string) {
	s[0] = "Go"
	// 这里 s 是个新的切片对象，
	// s 和 welcome 引用相同的数组
	fmt.Printf("s1 = %p\n", &s)
	s = append(s, "playground")
	fmt.Println(s)
	fmt.Printf("s2 = %p\n", &s)
}

func testVariadic() {
	welcome := []string{"hello", "world"}
	fmt.Printf("welcome %p\n", &welcome)
	change(welcome...)
	fmt.Println(welcome)
}

func main() {
	testVariadic()
}
