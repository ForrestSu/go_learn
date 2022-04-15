package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	b := []rune(s)
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, _ := Reverse(input)
	doubleRev, _ := Reverse(rev)
	var make = "ss"
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %s, %q\n", make, doubleRev)
}
