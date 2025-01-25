package utils

import (
	"fmt"
	"testing"
)

func Test_SliceToMatrix(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	grid := SliceToMatrix(list, 3)
	fmt.Println(grid)
}
