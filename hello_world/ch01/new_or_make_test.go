package main

import (
	"log"
	"testing"
)

func TestNewVsMake(t *testing.T) {
	// *[]int
	ptr := new([]int)
	log.Printf("%v %T, %p\n", ptr, ptr, ptr)

	// make(), 只适用于3种内建的引用类型：切片、map 和 channel
	obj := make([]int, 0) // []int
	log.Println(obj)
	log.Printf("%v %T, %p, %p, %T\n", obj, obj, obj, &obj, &obj)
}
