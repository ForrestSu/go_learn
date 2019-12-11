package main

import (
	"fmt"
	"unsafe"
)

func testBool() {

	a := true
	b := false
	fmt.Println("a:", a, "b:", b)

	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
}

func testInt() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d \n", a, unsafe.Sizeof(a)) //type and size of a
	fmt.Printf("type of b is %T, size of b is %d \n", b, unsafe.Sizeof(b)) //type and size of b
}

func testComplex() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)
}

// go don't do auto convert type
func testTypeAutoConvert() {
	// if define variable and don't use it, that an error!!
	i := 15   //int
	j := 17.8 //float64
	// sum := i + j //int + float64 not allowed
	fmt.Printf("j type is %T \n", j)

	sum := i + int(j)
	fmt.Println("sum :", sum) // 32
}

func explicitConvert() {
	i := 10
	// this statement will not work without explicit conversion
	var j = float64(i)
	fmt.Println("j: ", j)
}

func main() {

	testBool()
	testInt()
	testComplex()
	testTypeAutoConvert()
	explicitConvert()
}
