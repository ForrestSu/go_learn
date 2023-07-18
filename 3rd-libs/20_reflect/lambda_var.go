package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Doer struct {
	Data string
}

func (d *Doer) DoSomething(arg string) {

}

type Worker struct {
	DoSomething func(string)
}

func justHaveWorker(method func(string)) {
	var d Doer

	ptr := reflect.ValueOf(&method).UnsafePointer()
	env := reflect.NewAt(reflect.TypeOf(&d), unsafe.Pointer(*(*uintptr)(ptr)+(uintptr)(8)))

	fmt.Println(env.Type())
	a := env.Elem().Interface().(*Doer)
	fmt.Println(a.Data)
}

func main() {
	doer := Doer{Data: "hello"}
	method := doer.DoSomething
	justHaveWorker(method)
}
