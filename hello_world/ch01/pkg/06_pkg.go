package main

import (
	"fmt"
	"log"
	"pkg/rectangle"
)
/*
 * 1. package variables
 */

var rectLen, rectWidth float64 = 6, 7

/*
*2. init function to check if length and width are greater than zero
 */

func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

/**
最先初始化的是被导入的package，因此 rectangle package 先被初始化的。
接下来初始化的是 Package 级别变量 rectLen 和 rectWidth。
调用 init 函数
最后调用 main 函数
 */
func main() {
	// rectangle.init() //unexported name
	println("enter main...")
	area := rectangle.Area(rectLen, rectWidth)
	fmt.Println("area of rectangle = ", area)

}


/**
output:
rectangle package initialized
main package initialized
enter main...
area of rectangle =  42
 */