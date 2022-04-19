package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// visitorFunc visitor func
type visitorFunc func(shape Shape)

// Shape shape
type Shape interface {
	accept(fn visitorFunc)
}

// Circle circle
type Circle struct {
	Radius int
}

// Show show
func (c *Circle) accept(fn visitorFunc) {
	fn(c)
}

// Rectangle rectangle
type Rectangle struct {
	Width  int
	Height int
}

// implement Shape
func (r *Rectangle) accept(fn visitorFunc) {
	fn(r)
}

// JsonVisitor visitor
func JsonVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

// XmlVisitor visitor
func XmlVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
