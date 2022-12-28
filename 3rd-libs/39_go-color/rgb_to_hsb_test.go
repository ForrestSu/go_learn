package main

import (
	"testing"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func TestRGBA(t *testing.T) {
	rgb := colorful.Color{R: 0.313725, G: 0.478431, B: 0.721569}
	rgb.Hsv()
}

func TestRGBAWWW(t *testing.T) {
	// type Stu struct {
	// 	Name string `json:"name"`
	// }
	// stu := &Stu{}
	// json.Unmarshal([]byte(`{"name":"Bob"}`), stu)
	// t.Logf("%#v", stu)
	c, err := colorful.Hex("#993F2D")
	if err != nil {
		t.Log(err)
	}
	r, g, b := c.RGB255()
	t.Logf("%x %x %x", r, g, b)
}
