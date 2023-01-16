package go1_20

import (
	"errors"
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	err1 := errors.New("err1")
	err2 := errors.New("err2")

	err := errors.Join(err1, err2)
	fmt.Println(err)
}
