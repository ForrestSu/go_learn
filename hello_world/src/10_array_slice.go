package main

import "fmt"

func testArray() {
	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a)
	fmt.Println("length of a is",len(a))
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}
func testPassByValue() {
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
}

// 切片和python 切片类似，注意切片是引用，可以使用copy 进行深拷贝
func main() {
	testArray()
	testPassByValue()
	//testExpressionMatch()
	//testFallThrough()
}
