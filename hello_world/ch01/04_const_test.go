package main_test

import (
	"fmt"
	"testing"
)

func TestConst(t *testing.T) {
    const w = 10 // allow
    // a = 11 // reassignment not allowed
    fmt.Println("const a is ", w)

	fmt.Println("Hello, playground")
	//var a = math.Sqrt(4) //allowed
	//const b = math.Sqrt(4) //not allowed

	var name = "Sam"
	fmt.Printf("type %T value %v\n", name, name)

	const typedhello string = "Hello World"
}

/**
 go 是一种强类型语言
 */
func TestAliasType(t *testing.T) {
	//var defaultName = "Sam" //allowed
	// 这里是定义一个新类型，
	type myString string

	// 不可以用 等效的类型 直接赋值
	//var name myString = string("Sam") //not allowed

	//fmt.Println("name =", name)
	// 但是可以强制转换
	var customName myString = myString(string("Sam")) //allowed
	//customName = defaultName //not allowed

	fmt.Printf("=>> type %T value %v\n", customName, customName)
}

func TestDefaultType(t *testing.T) {
	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("i's type %T, f's type %T, c's type %T \n", i, f, c)
}

func TestVar(t *testing.T) {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)
}

// 数值常量可以在表达式中自由混合和匹配
