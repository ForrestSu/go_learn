package main

import (
	"fmt"

	bloom "github.com/bits-and-blooms/bloom/v3"
)

func main() {
	filter := bloom.NewWithEstimates(1000000, 0.01)

	// to add a string item, "Love"
	filter.Add([]byte("Love"))

	// Similarly, to test if "Love" is in bloom:
	if filter.Test([]byte("Love")) {
		fmt.Println("Love is in existed")
	} else {
		fmt.Println("Love is not exist")
	}
}
