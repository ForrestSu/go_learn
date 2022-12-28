package main

import (
	"fmt"
	"testing"
)

func handler() error {
	return fmt.Errorf("handler err is abandon")
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
	// err := Filter()
	// assert.Nil(t, err) // need err
	hello(2)
}

func hello(k int) {
	if k == 0 {
		fmt.Println("k is 0")
		return
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover")
		}
		fmt.Printf("defer \n")
	}()
	go func() {
		panic("panic")
	}()
}
