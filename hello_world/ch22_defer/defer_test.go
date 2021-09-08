package ch22_defer

import (
	"fmt"
	"testing"
)

func handler() error {
	return fmt.Errorf("handler err")
}

func Filter() (err error) {
	var verifyErr error
	defer func() {
		err = verifyErr
		fmt.Printf("err = %+v, verifyErr = %+v \n", err, verifyErr)
	}()

	verifyErr = nil
	return handler()
}

func TestHello(t *testing.T) {
	err := Filter()
	t.Logf("%+v", err)
}
