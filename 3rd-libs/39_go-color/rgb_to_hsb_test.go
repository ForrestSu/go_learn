package main

import (
	"testing"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func TestRGBA(t *testing.T) {
	rgb := colorful.Color{R: 0.313725, G: 0.478431, B: 0.721569}
	rgb.Hsv()
}
