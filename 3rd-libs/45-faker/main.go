package main

import (
	"fmt"

	"github.com/go-faker/faker/v4"
)

type SomeStruct struct {
	IsOnline bool
	Name     string
	Age      int32
	Tags     []string
}

func main() {
	a := SomeStruct{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}
