package main

import (
	"fmt"
	"testing"
)

// 4, 3, 2, 1
func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	fmt.Println("4")
}

type student struct {
	Name string
}

func TestIsSame(t *testing.T) {
	var stu = &student{Name: "Alice"}
	defer fmt.Printf("v1: %#v\n", stu)
	defer func() {
		fmt.Printf("v2: %#v\n", stu)
	}()
	stu.Name = "Bob"

	// output:
	// v2: &main.student{Name:"Bob"}
	// v1: &main.student{Name:"Bob"}
}

func TestIsDifferent(t *testing.T) {
	var stu = &student{Name: "Alice"}
	defer fmt.Printf("v1: %#v\n", stu)
	defer func() {
		fmt.Printf("v2: %#v\n", stu)
	}()
	stu = &student{Name: "Bob"}

	// output:
	// v2: &main.student{Name:"Bob"}
	// v1: &main.student{Name:"Alice"}
}
