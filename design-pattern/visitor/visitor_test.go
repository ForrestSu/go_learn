package main

import (
	"testing"
)

func TestCircle_show(t *testing.T) {
	c := Circle{Radius: 10}
	r := Rectangle{Width: 10, Height: 20}

	visitors := []Shape{&c, &r}
	for _, s := range visitors {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)
	}
}
